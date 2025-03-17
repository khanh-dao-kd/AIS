package consumer

import (
	"ais_service/internal/configs"
	"ais_service/internal/utils"
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"

	"cloud.google.com/go/pubsub"
	"go.uber.org/zap"
)

type Consumer interface {
	RegisterHandler(topicName string, subscriptionName string, handlerFunc HandlerFunc) error
	Start(ctx context.Context) error
}

type pubsubConsumer struct {
	pubsubClient              *pubsub.Client
	subscriptionToHandlerFunc map[string]HandlerFunc
	logger                    *zap.Logger
}

func NewPubSubConsumer(mqConfig configs.MQ, logger *zap.Logger) (Consumer, error) {
	ctx := context.Background()
	pubsubClient, err := pubsub.NewClient(ctx, mqConfig.ProjectID)
	if err != nil {
		return nil, fmt.Errorf("failed to create pubsub client: %w", err)
	}

	return &pubsubConsumer{
		pubsubClient:              pubsubClient,
		subscriptionToHandlerFunc: make(map[string]HandlerFunc),
		logger:                    logger,
	}, err
}

func (p *pubsubConsumer) RegisterHandler(topicName string, subscriptionName string, handlerFunc HandlerFunc) error {
	// p.logger.Info("Registering handler", zap.String("topic", topicName), zap.String("subscription", subscriptionName))

	// // Ensure the topic exists
	// topic := p.pubsubClient.Topic(topicName)
	// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// defer cancel()

	// topicExists, err := topic.Exists(ctx)
	// if err != nil {
	// 	p.logger.With(zap.Error(err)).Error("Failed to check topic existence")
	// 	return err
	// }

	// if !topicExists {
	// 	p.logger.Error("Topic does not exist", zap.String("topic_name", topicName))
	// 	return err
	// }

	// // Ensure the subscription exists
	// subscription := p.pubsubClient.Subscription(subscriptionName)
	// subExists, err := subscription.Exists(ctx)
	// if err != nil {
	// 	p.logger.With(zap.Error(err)).Error("Failed to check subscription existence")
	// 	return err
	// }

	// if !subExists {
	// 	p.logger.Warn("Subscription does not exist, creating it...", zap.String("subscription_name", subscriptionName))
	// 	_, err = p.pubsubClient.CreateSubscription(ctx, subscriptionName, pubsub.SubscriptionConfig{
	// 		Topic:       topic,
	// 		AckDeadline: 10 * time.Second,
	// 	})
	// 	if err != nil {
	// 		p.logger.With(zap.Error(err)).Error("Failed to create subscription")
	// 		return err
	// 	}
	// 	p.logger.Info("Subscription created successfully", zap.String("subscription_name", subscriptionName))
	// }

	p.subscriptionToHandlerFunc[subscriptionName] = handlerFunc
	return nil
}

func (p pubsubConsumer) Start(ctx context.Context) error {
	logger := utils.LoggerWithContext(ctx, p.logger)

	// Handle OS interrupts (graceful shutdown)
	exitSignalChannel := make(chan os.Signal, 1)
	signal.Notify(exitSignalChannel, os.Interrupt)

	var wg sync.WaitGroup

	for topicName, handlerFunc := range p.subscriptionToHandlerFunc {
		wg.Add(1)
		go func(topicName string, handlerFunc HandlerFunc) {
			defer wg.Done()
			sub := p.pubsubClient.Subscription(topicName)

			err := sub.Receive(ctx, func(ctx context.Context, msg *pubsub.Message) {

				// Call user-defined handler function
				if err := handlerFunc(ctx, topicName, msg.Data); err != nil {
					msg.Nack() // Mark message as not acknowledged (retries later)
					return
				}
				msg.Ack() // Acknowledge successful processing
			})

			if err != nil {
				// Handle fail consume message
				logger.
					With(zap.String("topic_name", topicName)).
					With(zap.Error(err)).
					Error("failed to consume message from queue")
			}
		}(topicName, handlerFunc)
	}

	// Wait for exit signal
	<-exitSignalChannel
	wg.Wait()

	return nil
}
