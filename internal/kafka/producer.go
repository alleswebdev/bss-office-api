package kafka

import (
	"context"
	"github.com/Shopify/sarama"
)

// NewSyncProducer создаёт экземпляр синхорнного продюсера для подключения к кафке
func NewSyncProducer(brokers []string) (sarama.SyncProducer, error) {
	config := sarama.NewConfig()
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true
	producer, err := sarama.NewSyncProducer(brokers, config)

	return producer, err
}

// SendMessage - отправляет сообщение в кафку
func SendMessage(ctx context.Context, producer sarama.SyncProducer, topic string, message []byte) error {
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.ByteEncoder(message),
	}

	_, _, err := producer.SendMessage(msg)

	return err
}
