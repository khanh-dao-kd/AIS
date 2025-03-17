package consumer

import (
	"ais_service/internal/dataaccess/mq/consumer"
	"ais_service/internal/dataaccess/mq/producer"
	"context"
	"encoding/json"
)

type ConsumerServer interface {
	Start(ctx context.Context) error
}

type consumerServer struct {
	accountCreatedHandler AccountCreatedHandler
	mqConsumer            consumer.Consumer
}

func NewConsumerServer(accountCreatedHandler AccountCreatedHandler, mqConsumer consumer.Consumer) ConsumerServer {
	return &consumerServer{
		accountCreatedHandler: accountCreatedHandler,
		mqConsumer:            mqConsumer,
	}
}

func (c consumerServer) Start(ctx context.Context) error {

	c.mqConsumer.RegisterHandler(
		producer.AISAccountTopic,
		AISAccountSubscription,
		func(ctx context.Context, topicName string, payload []byte) error {
			var event producer.AccountEvent
			if err := json.Unmarshal(payload, &event); err != nil {
				return err
			}
			return c.accountCreatedHandler.Handle(ctx, event)
		},
	)
	return c.mqConsumer.Start(ctx)
}
