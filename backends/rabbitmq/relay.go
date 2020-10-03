package rabbitmq

import (
	"context"
	"github.com/jhump/protoreflect/desc"
	"github.com/pkg/errors"
	"github.com/relistan/go-director"
	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"

	"github.com/batchcorp/plumber/api"
	"github.com/batchcorp/plumber/backends/rabbitmq/types"
	"github.com/batchcorp/plumber/cli"
	"github.com/batchcorp/plumber/relay"
)

type Relayer struct {
	Options        *cli.Options
	Channel        *amqp.Channel
	MsgDesc        *desc.MessageDescriptor
	RelayCh        chan interface{}
	log            *logrus.Entry
	Looper         *director.FreeLooper
	DefaultContext context.Context
}

func Relay(opts *cli.Options) error {
	if err := validateRelayOptions(opts); err != nil {
		return errors.Wrap(err, "unable to verify options")
	}

	// TODO: move this up the chain?
	ctx := context.Background()

	// Create new relayer instance (+ validate token & gRPC address)
	relayCfg := &relay.Config{
		Token:       opts.RelayToken,
		GRPCAddress: opts.RelayGRPCAddress,
		NumWorkers:  opts.RelayNumWorkers,
		Timeout:     opts.RelayGRPCTimeout,
		RelayCh:     make(chan interface{}, 1),
		DisableTLS:  opts.RelayGRPCDisableTLS,
	}

	grpcRelayer, err := relay.New(relayCfg)
	if err != nil {
		return errors.Wrap(err, "unable to create new gRPC relayer")
	}

	// Create new service
	channel, err := connect(opts)
	if err != nil {
		return errors.Wrap(err, "unable to create new RabbitMQ service")
	}

	// Launch HTTP server
	go func() {
		if err := api.Start(opts.RelayHTTPListenAddress, opts.Version); err != nil {
			logrus.Fatalf("unable to start API server: %s", err)
		}
	}()

	// Launch gRPC Relayer
	if err := grpcRelayer.StartWorkers(); err != nil {
		return errors.Wrap(err, "unable to start gRPC relay workers")
	}

	r := &Relayer{
		Channel:        channel,
		Options:        opts,
		RelayCh:        relayCfg.RelayCh,
		log:            logrus.WithField("pkg", "rabbitmq/relay"),
		Looper:         director.NewFreeLooper(director.FOREVER, make(chan error)),
		DefaultContext: ctx,
	}

	return r.Relay()
}

func validateRelayOptions(opts *cli.Options) error {
	if opts.Rabbit.RoutingKey == "" {
		return errors.New("You must specify a routing key")
	}
	if opts.Rabbit.ReadQueue == "" {
		return errors.New("You must specify a queue to read from")
	}
	if opts.Rabbit.Exchange == "" {
		return errors.New("You must specify an exchange")
	}
	return nil
}

func (r *Relayer) Relay() error {
	r.log.Infof("Relaying RabbitMQ messages from '%s' exchange -> '%s'",
		r.Options.Rabbit.Exchange, r.Options.RelayGRPCAddress)

	r.log.Infof("HTTP server listening on '%s'", r.Options.RelayHTTPListenAddress)

	msgChan, err := r.Channel.Consume(
		r.Options.Rabbit.ReadQueue,
		"",
		true,
		r.Options.Rabbit.ReadQueueExclusive,
		false,
		false,
		nil,
	)

	if err != nil {
		return errors.Wrap(err, "unable to create initial consume channel")
	}

	// Send message(s) to relayer
	r.Looper.Loop(func() error {
		select {
		case msg := <-msgChan:
			r.log.Debug("Writing RabbitMQ message to relay channel")

			r.RelayCh <- &types.RelayMessage{
				Value:   &msg,
				Options: &types.RelayMessageOptions{},
			}
		case <-r.DefaultContext.Done():
			r.log.Warning("received notice to quit loop")
			r.Looper.Quit()
		}
		return nil
	})

	return nil
}