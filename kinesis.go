package util

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kinesis"
)

type (
	// KinesisPublisher publishes given records to a kinesis stream
	KinesisPublisher interface {
		Publish([]*kinesis.PutRecordsRequestEntry) (*kinesis.PutRecordsOutput, error)
	}
	// KinesisPublisherFunc is a function adaptor to KinesisPublisher
	KinesisPublisherFunc func([]*kinesis.PutRecordsRequestEntry) (*kinesis.PutRecordsOutput, error)
)

// Publish is a function adaptor implementing the KinesisPublisher interface
func (k KinesisPublisherFunc) Publish(d []*kinesis.PutRecordsRequestEntry) (*kinesis.PutRecordsOutput, error) {
	return k(d)
}

// KPublisher implements KinesisPublisher
type KPublisher struct {
	svc       *kinesis.Kinesis
	outStream string
}

// NewKPublisher returns an Publisher for sending order items to kinesis
func NewKPublisher(outStream string) *KPublisher {
	k := new(KPublisher)
	k.outStream = outStream
	k.svc = kinesis.New(session.Must(session.NewSession()))
	return k
}

// Publish will publish the given records to a kinesis stream
func (k *KPublisher) Publish(records []*kinesis.PutRecordsRequestEntry) (*kinesis.PutRecordsOutput, error) {
	return k.svc.PutRecords(&kinesis.PutRecordsInput{
		Records:    records,
		StreamName: aws.String(k.outStream),
	})
}
