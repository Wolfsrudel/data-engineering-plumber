// Code generated by protoc-gen-go. DO NOT EDIT.
// source: opts/ps_opts_manage_connection.proto

package opts

import (
	fmt "fmt"
	args "github.com/batchcorp/plumber-schemas/build/go/protos/args"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type GetConnectionOptions struct {
	// @gotags: kong:"help='ID of the connection to get (leave empty to get all)'"
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" kong:"help='ID of the connection to get (leave empty to get all)'"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetConnectionOptions) Reset()         { *m = GetConnectionOptions{} }
func (m *GetConnectionOptions) String() string { return proto.CompactTextString(m) }
func (*GetConnectionOptions) ProtoMessage()    {}
func (*GetConnectionOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_135d08c7ffd8961f, []int{0}
}

func (m *GetConnectionOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetConnectionOptions.Unmarshal(m, b)
}
func (m *GetConnectionOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetConnectionOptions.Marshal(b, m, deterministic)
}
func (m *GetConnectionOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetConnectionOptions.Merge(m, src)
}
func (m *GetConnectionOptions) XXX_Size() int {
	return xxx_messageInfo_GetConnectionOptions.Size(m)
}
func (m *GetConnectionOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_GetConnectionOptions.DiscardUnknown(m)
}

var xxx_messageInfo_GetConnectionOptions proto.InternalMessageInfo

func (m *GetConnectionOptions) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type CreateConnectionOptions struct {
	// @gotags: kong:"help='Friendly name for the connection', required"
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty" kong:"help='Friendly name for the connection', required"`
	// @gotags: kong:"help='Optional notes to associate with the connection'"
	Notes string `protobuf:"bytes,2,opt,name=notes,proto3" json:"notes,omitempty" kong:"help='Optional notes to associate with the connection'"`
	// @gotags: kong:"cmd,help='Apache Kafka'"
	Kafka *args.KafkaConn `protobuf:"bytes,100,opt,name=kafka,proto3" json:"kafka,omitempty" kong:"cmd,help='Apache Kafka'"`
	// @gotags: kong:"cmd,help='ActiveMQ'"
	ActiveMq *args.ActiveMQConn `protobuf:"bytes,101,opt,name=active_mq,json=activeMq,proto3" json:"active_mq,omitempty" kong:"cmd,help='ActiveMQ'"`
	// @gotags: kong:"cmd,help='AWS Simple Queue Service'"
	AwsSqs *args.AWSSQSConn `protobuf:"bytes,102,opt,name=aws_sqs,json=awsSqs,proto3" json:"aws_sqs,omitempty" kong:"cmd,help='AWS Simple Queue Service'"`
	// @gotags: kong:"cmd,help='AWS Simple Notification Service'"
	AwsSns *args.AWSSNSConn `protobuf:"bytes,103,opt,name=aws_sns,json=awsSns,proto3" json:"aws_sns,omitempty" kong:"cmd,help='AWS Simple Notification Service'"`
	// @gotags: kong:"cmd,help='Mongo (CDC)'"
	Mongo *args.MongoConn `protobuf:"bytes,104,opt,name=mongo,proto3" json:"mongo,omitempty" kong:"cmd,help='Mongo (CDC)'"`
	// @gotags: kong:"cmd,help='NATS PubSub'"
	Nats *args.NatsConn `protobuf:"bytes,105,opt,name=nats,proto3" json:"nats,omitempty" kong:"cmd,help='NATS PubSub'"`
	// @gotags: kong:"cmd,help='NATS Streaming (deprecated)'"
	NatsStreaming *args.NatsStreamingConn `protobuf:"bytes,106,opt,name=nats_streaming,json=natsStreaming,proto3" json:"nats_streaming,omitempty" kong:"cmd,help='NATS Streaming (deprecated)'"`
	// @gotags: kong:"cmd,help='NSQ'"
	Nsq *args.NSQConn `protobuf:"bytes,107,opt,name=nsq,proto3" json:"nsq,omitempty" kong:"cmd,help='NSQ'"`
	// @gotags: kong:"cmd,help='PostgreSQL (CDC)'"
	Postgres *args.PostgresConn `protobuf:"bytes,108,opt,name=postgres,proto3" json:"postgres,omitempty" kong:"cmd,help='PostgreSQL (CDC)'"`
	// @gotags: kong:"cmd,help='Pulsar'"
	Pulsar *args.PulsarConn `protobuf:"bytes,109,opt,name=pulsar,proto3" json:"pulsar,omitempty" kong:"cmd,help='Pulsar'"`
	// @gotags: kong:"cmd,help='Rabbit'"
	Rabbit *args.RabbitConn `protobuf:"bytes,110,opt,name=rabbit,proto3" json:"rabbit,omitempty" kong:"cmd,help='Rabbit'"`
	// @gotags: kong:"cmd,help='Rabbit Streams'"
	RabbitStreams *args.RabbitStreamsConn `protobuf:"bytes,111,opt,name=rabbit_streams,json=rabbitStreams,proto3" json:"rabbit_streams,omitempty" kong:"cmd,help='Rabbit Streams'"`
	// @gotags: kong:"cmd,help='Redis PubSub'"
	RedisPubsub *args.RedisPubSubConn `protobuf:"bytes,112,opt,name=redis_pubsub,json=redisPubsub,proto3" json:"redis_pubsub,omitempty" kong:"cmd,help='Redis PubSub'"`
	// @gotags: kong:"cmd,help='Redis Streams'"
	RedisStreams *args.RedisStreamsConn `protobuf:"bytes,113,opt,name=redis_streams,json=redisStreams,proto3" json:"redis_streams,omitempty" kong:"cmd,help='Redis Streams'"`
	// @gotags: kong:"cmd,help='Azure Event Hub'"
	AzureEventHub *args.AzureEventHubConn `protobuf:"bytes,114,opt,name=azure_event_hub,json=azureEventHub,proto3" json:"azure_event_hub,omitempty" kong:"cmd,help='Azure Event Hub'"`
	// @gotags: kong:"cmd,help='Azure Service Bus'"
	AzureServiceBus *args.AzureServiceBusConn `protobuf:"bytes,115,opt,name=azure_service_bus,json=azureServiceBus,proto3" json:"azure_service_bus,omitempty" kong:"cmd,help='Azure Service Bus'"`
	// @gotags: kong:"cmd,help='MQTT'"
	Mqtt *args.MQTTConn `protobuf:"bytes,116,opt,name=mqtt,proto3" json:"mqtt,omitempty" kong:"cmd,help='MQTT'"`
	// @gotags: kong:"cmd,help='KubeMQ Queue'"
	KubemqQueue *args.KubeMQQueueConn `protobuf:"bytes,117,opt,name=kubemq_queue,json=kubemqQueue,proto3" json:"kubemq_queue,omitempty" kong:"cmd,help='KubeMQ Queue'"`
	// @gotags: kong:"cmd,help='Google Cloud Pub/Sub'"
	GcpPubsub *args.GCPPubSubConn `protobuf:"bytes,118,opt,name=gcp_pubsub,json=gcpPubsub,proto3" json:"gcp_pubsub,omitempty" kong:"cmd,help='Google Cloud Pub/Sub'"`
	// @gotags: kong:"cmd,help='NATS JetStream'"
	NatsJetstream *args.NatsJetstreamConn `protobuf:"bytes,119,opt,name=nats_jetstream,json=natsJetstream,proto3" json:"nats_jetstream,omitempty" kong:"cmd,help='NATS JetStream'"`
	// @gotags: kong:"cmd,help='AWS Kinesis'"
	AwsKinesis           *args.AWSKinesisConn `protobuf:"bytes,120,opt,name=aws_kinesis,json=awsKinesis,proto3" json:"aws_kinesis,omitempty" kong:"cmd,help='AWS Kinesis'"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *CreateConnectionOptions) Reset()         { *m = CreateConnectionOptions{} }
func (m *CreateConnectionOptions) String() string { return proto.CompactTextString(m) }
func (*CreateConnectionOptions) ProtoMessage()    {}
func (*CreateConnectionOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_135d08c7ffd8961f, []int{1}
}

func (m *CreateConnectionOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateConnectionOptions.Unmarshal(m, b)
}
func (m *CreateConnectionOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateConnectionOptions.Marshal(b, m, deterministic)
}
func (m *CreateConnectionOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateConnectionOptions.Merge(m, src)
}
func (m *CreateConnectionOptions) XXX_Size() int {
	return xxx_messageInfo_CreateConnectionOptions.Size(m)
}
func (m *CreateConnectionOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateConnectionOptions.DiscardUnknown(m)
}

var xxx_messageInfo_CreateConnectionOptions proto.InternalMessageInfo

func (m *CreateConnectionOptions) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateConnectionOptions) GetNotes() string {
	if m != nil {
		return m.Notes
	}
	return ""
}

func (m *CreateConnectionOptions) GetKafka() *args.KafkaConn {
	if m != nil {
		return m.Kafka
	}
	return nil
}

func (m *CreateConnectionOptions) GetActiveMq() *args.ActiveMQConn {
	if m != nil {
		return m.ActiveMq
	}
	return nil
}

func (m *CreateConnectionOptions) GetAwsSqs() *args.AWSSQSConn {
	if m != nil {
		return m.AwsSqs
	}
	return nil
}

func (m *CreateConnectionOptions) GetAwsSns() *args.AWSSNSConn {
	if m != nil {
		return m.AwsSns
	}
	return nil
}

func (m *CreateConnectionOptions) GetMongo() *args.MongoConn {
	if m != nil {
		return m.Mongo
	}
	return nil
}

func (m *CreateConnectionOptions) GetNats() *args.NatsConn {
	if m != nil {
		return m.Nats
	}
	return nil
}

func (m *CreateConnectionOptions) GetNatsStreaming() *args.NatsStreamingConn {
	if m != nil {
		return m.NatsStreaming
	}
	return nil
}

func (m *CreateConnectionOptions) GetNsq() *args.NSQConn {
	if m != nil {
		return m.Nsq
	}
	return nil
}

func (m *CreateConnectionOptions) GetPostgres() *args.PostgresConn {
	if m != nil {
		return m.Postgres
	}
	return nil
}

func (m *CreateConnectionOptions) GetPulsar() *args.PulsarConn {
	if m != nil {
		return m.Pulsar
	}
	return nil
}

func (m *CreateConnectionOptions) GetRabbit() *args.RabbitConn {
	if m != nil {
		return m.Rabbit
	}
	return nil
}

func (m *CreateConnectionOptions) GetRabbitStreams() *args.RabbitStreamsConn {
	if m != nil {
		return m.RabbitStreams
	}
	return nil
}

func (m *CreateConnectionOptions) GetRedisPubsub() *args.RedisPubSubConn {
	if m != nil {
		return m.RedisPubsub
	}
	return nil
}

func (m *CreateConnectionOptions) GetRedisStreams() *args.RedisStreamsConn {
	if m != nil {
		return m.RedisStreams
	}
	return nil
}

func (m *CreateConnectionOptions) GetAzureEventHub() *args.AzureEventHubConn {
	if m != nil {
		return m.AzureEventHub
	}
	return nil
}

func (m *CreateConnectionOptions) GetAzureServiceBus() *args.AzureServiceBusConn {
	if m != nil {
		return m.AzureServiceBus
	}
	return nil
}

func (m *CreateConnectionOptions) GetMqtt() *args.MQTTConn {
	if m != nil {
		return m.Mqtt
	}
	return nil
}

func (m *CreateConnectionOptions) GetKubemqQueue() *args.KubeMQQueueConn {
	if m != nil {
		return m.KubemqQueue
	}
	return nil
}

func (m *CreateConnectionOptions) GetGcpPubsub() *args.GCPPubSubConn {
	if m != nil {
		return m.GcpPubsub
	}
	return nil
}

func (m *CreateConnectionOptions) GetNatsJetstream() *args.NatsJetstreamConn {
	if m != nil {
		return m.NatsJetstream
	}
	return nil
}

func (m *CreateConnectionOptions) GetAwsKinesis() *args.AWSKinesisConn {
	if m != nil {
		return m.AwsKinesis
	}
	return nil
}

type DeleteConnectionOptions struct {
	// @gotags: kong:"help='ID of the connection to delete',required=true"
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" kong:"help='ID of the connection to delete',required=true"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteConnectionOptions) Reset()         { *m = DeleteConnectionOptions{} }
func (m *DeleteConnectionOptions) String() string { return proto.CompactTextString(m) }
func (*DeleteConnectionOptions) ProtoMessage()    {}
func (*DeleteConnectionOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_135d08c7ffd8961f, []int{2}
}

func (m *DeleteConnectionOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteConnectionOptions.Unmarshal(m, b)
}
func (m *DeleteConnectionOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteConnectionOptions.Marshal(b, m, deterministic)
}
func (m *DeleteConnectionOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteConnectionOptions.Merge(m, src)
}
func (m *DeleteConnectionOptions) XXX_Size() int {
	return xxx_messageInfo_DeleteConnectionOptions.Size(m)
}
func (m *DeleteConnectionOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteConnectionOptions.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteConnectionOptions proto.InternalMessageInfo

func (m *DeleteConnectionOptions) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func init() {
	proto.RegisterType((*GetConnectionOptions)(nil), "protos.opts.GetConnectionOptions")
	proto.RegisterType((*CreateConnectionOptions)(nil), "protos.opts.CreateConnectionOptions")
	proto.RegisterType((*DeleteConnectionOptions)(nil), "protos.opts.DeleteConnectionOptions")
}

func init() {
	proto.RegisterFile("opts/ps_opts_manage_connection.proto", fileDescriptor_135d08c7ffd8961f)
}

var fileDescriptor_135d08c7ffd8961f = []byte{
	// 795 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x96, 0xdf, 0x4f, 0xdb, 0x3a,
	0x14, 0xc7, 0x05, 0xf7, 0xc2, 0xa5, 0xee, 0x2d, 0x57, 0xd7, 0x62, 0xd4, 0x94, 0xc1, 0xba, 0x0a,
	0x21, 0x90, 0xb6, 0x66, 0xda, 0xb4, 0x49, 0x68, 0x93, 0x26, 0x60, 0x8c, 0x69, 0xac, 0xac, 0x25,
	0x48, 0x93, 0xf6, 0x12, 0x39, 0xa9, 0x49, 0x43, 0x1b, 0xe7, 0x87, 0xed, 0x76, 0xda, 0xdf, 0xbe,
	0x87, 0xc9, 0x76, 0x52, 0x62, 0x6a, 0xf6, 0x42, 0x88, 0xbf, 0x9f, 0xef, 0xc9, 0xf1, 0xc9, 0x39,
	0x4e, 0xc1, 0x5e, 0x92, 0x72, 0xe6, 0xa4, 0xcc, 0x93, 0x57, 0x2f, 0xc6, 0x14, 0x87, 0xc4, 0x0b,
	0x12, 0x4a, 0x49, 0xc0, 0xa3, 0x84, 0x76, 0xd3, 0x3c, 0xe1, 0x09, 0xac, 0xab, 0x0b, 0xeb, 0x4a,
	0xa8, 0xb5, 0x8d, 0xf3, 0x50, 0x59, 0xe4, 0xd5, 0xc3, 0x01, 0x8f, 0xa6, 0x24, 0xce, 0x34, 0xd9,
	0xda, 0x35, 0xc5, 0x19, 0xf3, 0xc6, 0x11, 0x25, 0x2c, 0x62, 0x85, 0xde, 0x5a, 0xd0, 0x19, 0xfd,
	0x83, 0x96, 0x95, 0x5a, 0xc7, 0xd4, 0x7e, 0x8a, 0x9c, 0x78, 0x64, 0x4a, 0x28, 0xf7, 0x46, 0xc2,
	0x2f, 0x98, 0x3d, 0x0b, 0xc3, 0x48, 0x3e, 0x8d, 0x02, 0xe2, 0xf9, 0xa2, 0x8c, 0xb4, 0x63, 0x50,
	0x61, 0x90, 0x7a, 0xa9, 0xf0, 0xd9, 0x3c, 0x08, 0x32, 0xe4, 0x31, 0xbe, 0x19, 0xe3, 0x42, 0x79,
	0x62, 0x2a, 0xc2, 0x27, 0x71, 0xe6, 0x65, 0x82, 0x08, 0x62, 0xb5, 0xc6, 0x09, 0x0d, 0x93, 0x42,
	0x69, 0x9a, 0x4a, 0xc6, 0xb9, 0x55, 0xa0, 0x98, 0x97, 0x59, 0x3e, 0x5d, 0x10, 0xbc, 0x5b, 0xc2,
	0x19, 0xcf, 0x09, 0x8e, 0x1f, 0x46, 0xb4, 0x1e, 0xd1, 0xb0, 0x40, 0x36, 0x4d, 0x84, 0x95, 0x6f,
	0xc9, 0x7c, 0x85, 0x69, 0xc2, 0x78, 0x98, 0x93, 0xf2, 0xd1, 0x5b, 0xa6, 0x28, 0x26, 0x0c, 0xe7,
	0x56, 0x29, 0xc7, 0xbe, 0x1f, 0x71, 0x6b, 0x36, 0x5a, 0x2a, 0xf2, 0x61, 0xd6, 0x02, 0xe6, 0x64,
	0x18, 0x31, 0xb3, 0xf6, 0x6d, 0x0b, 0x60, 0x84, 0xe8, 0xec, 0x83, 0x8d, 0x73, 0xc2, 0x4f, 0xe7,
	0xfd, 0xf9, 0x35, 0x95, 0x7f, 0x19, 0x5c, 0x07, 0xcb, 0xd1, 0x10, 0x2d, 0xb5, 0x97, 0x0e, 0x6a,
	0x57, 0xcb, 0xd1, 0xb0, 0xf3, 0xab, 0x06, 0x9a, 0xa7, 0x39, 0xc1, 0x9c, 0x2c, 0xb2, 0x10, 0xfc,
	0x4d, 0x71, 0x4c, 0x0a, 0x5a, 0xfd, 0x0f, 0x37, 0xc0, 0x0a, 0x4d, 0x38, 0x61, 0x68, 0x59, 0x2d,
	0xea, 0x1b, 0xf8, 0x0c, 0xac, 0xa8, 0x06, 0x40, 0xc3, 0xf6, 0xd2, 0x41, 0xfd, 0xe5, 0x66, 0xb7,
	0x18, 0x03, 0x99, 0x5e, 0xf7, 0x42, 0x2a, 0x32, 0xfa, 0x95, 0x86, 0xe0, 0x1b, 0x50, 0xd3, 0xc3,
	0xe0, 0xc5, 0x19, 0x22, 0xca, 0xb1, 0x65, 0x38, 0x8e, 0x95, 0xda, 0x1b, 0x28, 0xd3, 0x9a, 0x66,
	0x7b, 0x19, 0x7c, 0x01, 0xfe, 0x29, 0x7a, 0x1d, 0xdd, 0x28, 0x57, 0xd3, 0x74, 0x7d, 0x73, 0xdd,
	0x81, 0xab, 0x3c, 0xab, 0x78, 0xc6, 0xdc, 0x8c, 0xcd, 0x1d, 0x94, 0xa1, 0xf0, 0x01, 0xc7, 0x65,
	0xc5, 0x41, 0xd5, 0x4e, 0x54, 0x3f, 0xa2, 0x91, 0x65, 0x27, 0x3d, 0xa9, 0xe8, 0x9d, 0x28, 0x08,
	0x1e, 0xca, 0x0a, 0x71, 0x86, 0x22, 0x05, 0x3f, 0x32, 0xe0, 0x4b, 0xcc, 0x99, 0x62, 0x15, 0x02,
	0xcf, 0xc0, 0xba, 0xd9, 0x79, 0xe8, 0x56, 0x99, 0x76, 0x17, 0x4c, 0x6e, 0x49, 0x28, 0x77, 0x83,
	0x56, 0x97, 0xe0, 0x3e, 0xf8, 0x8b, 0xb2, 0x0c, 0x8d, 0x95, 0x77, 0xc3, 0xf4, 0xba, 0xba, 0x60,
	0x12, 0x80, 0xaf, 0xc1, 0x5a, 0xd9, 0xad, 0x68, 0x62, 0x29, 0x71, 0xbf, 0x10, 0x75, 0x89, 0x4b,
	0x14, 0x3a, 0x60, 0x55, 0xf7, 0x31, 0x8a, 0x2d, 0xf5, 0xea, 0x2b, 0x49, 0xd7, 0x4b, 0x63, 0xd2,
	0xa0, 0x5b, 0x18, 0x51, 0x8b, 0xe1, 0x4a, 0x49, 0xda, 0xa0, 0x31, 0x59, 0x07, 0xb3, 0xe7, 0x51,
	0x62, 0xa9, 0x83, 0x36, 0xea, 0x6d, 0xeb, 0x1c, 0x1b, 0x79, 0x75, 0x09, 0xbe, 0x07, 0xff, 0x56,
	0xe7, 0x02, 0xa5, 0x2a, 0xc8, 0x63, 0x33, 0x88, 0x04, 0xfa, 0xc2, 0x77, 0x85, 0xaf, 0x42, 0xd4,
	0xf3, 0x62, 0x81, 0x09, 0x1f, 0x9e, 0x80, 0x86, 0x31, 0x37, 0x28, 0x53, 0x11, 0x76, 0x16, 0x23,
	0x54, 0xb3, 0xd0, 0x0f, 0x2d, 0x93, 0xf8, 0x08, 0xfe, 0xbb, 0x77, 0xc0, 0xa2, 0xdc, 0xb2, 0x99,
	0x63, 0xc9, 0x9c, 0x49, 0xe4, 0x53, 0x91, 0x49, 0x03, 0x57, 0x97, 0xe0, 0x17, 0xf0, 0xff, 0xc2,
	0x21, 0x8c, 0x98, 0x8a, 0xd4, 0x5e, 0x8c, 0xe4, 0x6a, 0xe8, 0x44, 0xe8, 0x94, 0x74, 0x0a, 0x77,
	0x8b, 0xb2, 0x29, 0xe5, 0xc1, 0x89, 0xb8, 0xa5, 0x29, 0x7b, 0x83, 0xeb, 0x6b, 0xdd, 0x94, 0x12,
	0x91, 0x55, 0xac, 0x1e, 0xcf, 0x48, 0x58, 0xaa, 0x78, 0x21, 0x7c, 0xd2, 0x1b, 0x0c, 0xa4, 0xae,
	0xab, 0xa8, 0x1d, 0x6a, 0x01, 0x1e, 0x01, 0x70, 0xf7, 0x61, 0x40, 0x53, 0x65, 0x6f, 0x19, 0xf6,
	0xf3, 0xd3, 0x7e, 0xe5, 0x15, 0xd4, 0xc2, 0x20, 0x2d, 0x5e, 0x40, 0x39, 0x10, 0xf3, 0xd3, 0x1a,
	0xcd, 0x1e, 0x18, 0x88, 0xcf, 0x25, 0x71, 0x37, 0x10, 0xf3, 0x25, 0xf8, 0x0e, 0xd4, 0x2b, 0x1f,
	0x4f, 0xf4, 0x43, 0xc5, 0xd8, 0xbe, 0x3f, 0xe6, 0x17, 0x5a, 0x56, 0x01, 0x00, 0x9e, 0xb1, 0xe2,
	0xbe, 0x73, 0x08, 0x9a, 0x1f, 0xc8, 0x84, 0xd8, 0x4e, 0xbf, 0x7b, 0x27, 0xe5, 0xc9, 0xdb, 0xef,
	0x47, 0x61, 0xc4, 0xe5, 0x47, 0x34, 0x48, 0x62, 0xc7, 0xc7, 0x3c, 0x18, 0x05, 0x49, 0x9e, 0x3a,
	0xe9, 0x44, 0xc4, 0x3e, 0xc9, 0x9f, 0xb3, 0x60, 0x44, 0x62, 0xcc, 0x1c, 0x5f, 0x44, 0x93, 0xa1,
	0x13, 0x26, 0x8e, 0x4e, 0xc1, 0x91, 0x3f, 0x05, 0xfc, 0x55, 0x75, 0xf3, 0xea, 0x77, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x14, 0xce, 0x90, 0x3b, 0x46, 0x08, 0x00, 0x00,
}