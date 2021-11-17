package kubemq_queue

import (
	"context"

	queuesStream "github.com/kubemq-io/kubemq-go/queues_stream"
	"github.com/pkg/errors"

	"github.com/batchcorp/plumber-schemas/build/go/protos/opts"
	"github.com/batchcorp/plumber-schemas/build/go/protos/records"

	"github.com/batchcorp/plumber/backends/kubemq-queue/types"
	"github.com/batchcorp/plumber/prometheus"
)

func (k *KubeMQ) Relay(ctx context.Context, relayOpts *opts.RelayOptions, relayCh chan interface{}, errorCh chan *records.ErrorRecord) error {
	if err := validateRelayOptions(relayOpts); err != nil {
		return errors.Wrap(err, "unable to validate relay options")
	}

	for {
		select {
		case <-ctx.Done():
			k.log.Info("Received shutdown signal, exiting relayer")
			return nil
		default:
			// NOOP
		}

		response, err := k.client.Poll(context.Background(),
			queuesStream.NewPollRequest().
				SetChannel(relayOpts.KubemqQueue.Args.QueueName).
				SetMaxItems(1). // TODO: flag?
				SetAutoAck(false).
				SetWaitTimeout(DefaultReadTimeout))
		if err != nil {
			return err
		}

		if !response.HasMessages() {
			continue
		}

		if err := response.AckAll(); err != nil {
			return errors.Wrap(err, "unable to acknowledge message(s)")
		}

		for _, msg := range response.Messages {
			relayCh <- &types.RelayMessage{
				Value: msg,
			}

			prometheus.Incr("kubemq-queue-relay-consumer", 1)
		}

	}

	return nil
}

func validateRelayOptions(relayOpts *opts.RelayOptions) error {
	if relayOpts == nil {
		return errors.New("relay options cannot be nil")
	}

	if relayOpts.KubemqQueue == nil {
		return errors.New("backend group options cannot be nil")
	}

	if relayOpts.KubemqQueue.Args == nil {
		return errors.New("backend arg options cannot be nil")
	}

	if relayOpts.KubemqQueue.Args.QueueName == "" {
		return errors.New("queue name cannot be empty")
	}

	return nil
}