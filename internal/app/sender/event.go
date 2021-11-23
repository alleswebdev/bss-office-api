// Package sender предназначен для отправки событий в брокер сообщений
package sender

import (
	"context"
	"github.com/Shopify/sarama"
	"github.com/ozonmp/bss-office-api/internal/kafka"
	"github.com/ozonmp/bss-office-api/internal/model"
	"github.com/pkg/errors"
	"google.golang.org/protobuf/proto"
)

// EventSender interface
type EventSender interface {
	Send(ctx context.Context, office *model.OfficeEvent) error
}

type eventSender struct {
	producer sarama.SyncProducer
	topic    string
}

func NewEventSender(brokers []string, topic string) (*eventSender, error) {
	syncProducer, err := kafka.NewSyncProducer(brokers)

	if err != nil {
		return nil, err
	}

	sender := &eventSender{producer: syncProducer, topic: topic}

	return sender, nil
}

func (s *eventSender) Send(ctx context.Context, officeEvent *model.OfficeEvent) error {
	pb := model.ConvertBssOfficeEventToPb(officeEvent)

	msg, err := proto.Marshal(pb)
	if err != nil {
		return errors.Wrap(err, "EventSender.Marshal()")
	}

	err = kafka.SendMessage(ctx, s.producer, s.topic, msg)

	if err != nil {
		return errors.Wrap(err, "EventSender.SendMessage()")
	}

	return nil
}
