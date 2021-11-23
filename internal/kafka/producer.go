package kafka

import (
	"context"
	"github.com/Shopify/sarama"
	"github.com/ozonmp/bss-office-api/internal/logger"
)

func NewSyncProducer(brokers []string) (sarama.SyncProducer, error) {
	config := sarama.NewConfig()
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true
	producer, err := sarama.NewSyncProducer(brokers, config)

	return producer, err
}

func SendMessage(ctx context.Context, producer sarama.SyncProducer, topic string, message []byte) error {
	msg := &sarama.ProducerMessage{
		Topic:     topic,
		Partition: -1,
		Value:     sarama.ByteEncoder(message),
	}

	part, offset, err := producer.SendMessage(msg)

	logger.DebugKV(ctx, "Send msg", "msg", msg, "partition", part, "offset", offset)

	return err
}
