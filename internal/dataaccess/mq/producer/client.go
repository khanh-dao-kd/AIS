package producer

import (
	"ais_service/internal/configs"
	"context"
	"fmt"

	"cloud.google.com/go/pubsub"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Client interface {
	Produce(ctx context.Context, topicName string, payload []byte) error
}

type client struct {
	pubsubClient *pubsub.Client
}

func NewClient(ctx context.Context, mqConfig configs.MQ) (Client, error) {
	pubsubClient, err := pubsub.NewClient(ctx, mqConfig.ProjectID)
	if err != nil {
		return nil, fmt.Errorf("failed to create pubsub client: %w", err)
	}

	return &client{
		pubsubClient: pubsubClient,
	}, nil
}

func (c client) Produce(ctx context.Context, topicName string, payload []byte) error {
	topic := c.pubsubClient.Topic(topicName)
	defer topic.Stop()

	result := topic.Publish(ctx, &pubsub.Message{
		Data: payload,
	})

	// Wait for the result to be acknowledged
	_, err := result.Get(ctx)
	if err != nil {
		return status.Error(codes.Internal, "failed to publish message")
	}
	return nil
}
