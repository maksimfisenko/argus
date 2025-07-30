package kafka

import (
	"time"

	"github.com/segmentio/kafka-go"
	"golang.org/x/net/context"
)

type Producer struct {
	writer *kafka.Writer
}

func NewProducer(brokers []string, topic string) *Producer {
	return &Producer{
		writer: kafka.NewWriter(kafka.WriterConfig{
			Brokers:  brokers,
			Topic:    topic,
			Balancer: &kafka.LeastBytes{},
		}),
	}
}

func (p *Producer) Publish(ctx context.Context, data []byte) error {
	msg := kafka.Message{
		Key:   []byte(time.Now().Format(time.RFC3339Nano)),
		Value: data,
	}
	return p.writer.WriteMessages(ctx, msg)
}

func (p *Producer) Close() error {
	return p.writer.Close()
}
