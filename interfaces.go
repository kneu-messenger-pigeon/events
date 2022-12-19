package events

import (
	"context"
	"github.com/segmentio/kafka-go"
)

type WriterInterface interface {
	WriteMessages(ctx context.Context, msgs ...kafka.Message) error
	Close() error
}

type ReaderInterface interface {
	ReadMessage(context.Context) (kafka.Message, error)
	FetchMessage(context.Context) (kafka.Message, error)
	CommitMessages(context.Context, ...kafka.Message) error
	Close() error
}
