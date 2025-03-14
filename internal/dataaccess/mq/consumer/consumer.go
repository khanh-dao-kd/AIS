package consumer

import (
	"ais_service/internal/configs"
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"

	"cloud.google.com/go/pubsub"
)

type Consumer interface {
	RegisterHandler(topicName string, handlerFunc HandlerFunc)
	Start(ctx context.Context) error
}

type pubsubConsumer struct {
	pubsubClient              *pubsub.Client
	subscriptionToHandlerFunc map[string]HandlerFunc
}

func NewPubSubConsumer(mqConfig configs.MQ) (Consumer, error) {
	ctx := context.Background()
	pubsubClient, err := pubsub.NewClient(ctx, mqConfig.ProjectID)
	if err != nil {
		return nil, fmt.Errorf("failed to create pubsub client: %w", err)
	}

	return &pubsubConsumer{
		pubsubClient:              pubsubClient,
		subscriptionToHandlerFunc: make(map[string]HandlerFunc),
	}, err
}

func (p pubsubConsumer) RegisterHandler(topicName string, handlerFunc HandlerFunc) {
	p.subscriptionToHandlerFunc[topicName] = handlerFunc
}

func (p pubsubConsumer) Start(ctx context.Context) error {
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
			}
		}(topicName, handlerFunc)
	}

	// Wait for exit signal
	<-exitSignalChannel
	wg.Wait()

	return nil
}
