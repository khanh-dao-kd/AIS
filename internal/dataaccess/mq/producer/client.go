package producer

import (
	"ais_service/internal/configs"
	"ais_service/internal/utils"
	"context"
	"fmt"

	"cloud.google.com/go/pubsub"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Client interface {
	Produce(ctx context.Context, topicName string, payload []byte) error
}

type client struct {
	pubsubClient *pubsub.Client
	logger       *zap.Logger
}

func NewClient(mqConfig configs.MQ, logger *zap.Logger) (Client, error) {
	ctx := context.Background()
	pubsubClient, err := pubsub.NewClient(ctx, mqConfig.ProjectID)
	if err != nil {
		return nil, fmt.Errorf("failed to create pubsub client: %w", err)
	}

	return &client{
		pubsubClient: pubsubClient,
		logger:       logger,
	}, nil
}

func (c client) Produce(ctx context.Context, topicName string, payload []byte) error {
	logger := utils.LoggerWithContext(ctx, c.logger).
		With(zap.String("topic_name", topicName)).
		With(zap.ByteString("payload", payload))

	// c.CreateTopic(ctx, topicName)

	topic := c.pubsubClient.Topic(topicName)
	defer topic.Stop()

	result := topic.Publish(ctx, &pubsub.Message{
		Data: payload,
	})

	// Wait for the result to be acknowledged
	_, err := result.Get(ctx)
	if err != nil {
		logger.With(zap.Error(err)).Error("failed to publish message")
		return status.Error(codes.Internal, "failed to publish message")
	}
	return nil
}

func (c client) CreateTopic(ctx context.Context, topicName string) error {
	logger := utils.LoggerWithContext(ctx, c.logger).With(zap.String("topic_name", topicName))

	topic := c.pubsubClient.Topic(topicName)

	// Check if the topic exists
	exists, err := topic.Exists(ctx)
	if err != nil {
		logger.With(zap.Error(err)).Error("Failed to check topic existence")
		return status.Error(codes.Internal, "failed to check topic existence")
	}

	// Create the topic if it does not exist
	if !exists {
		logger.Warn("Topic does not exist, creating it...")
		topic, err = c.pubsubClient.CreateTopic(ctx, topicName)
		if err != nil {
			logger.With(zap.Error(err)).Error("Failed to create topic")
			return status.Error(codes.Internal, "failed to create topic")
		}
		logger.Info("Topic created successfully", zap.String("topic_name", topic.ID()))
	}
	return nil
}
