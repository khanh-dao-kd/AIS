package server

import (
	"ais_service/internal/configs"
	"ais_service/internal/dataaccess/mq/producer"
	"ais_service/internal/handler/consumer"
	"ais_service/internal/handler/grpc"
	"ais_service/internal/handler/http"
	"ais_service/internal/utils"
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"cloud.google.com/go/pubsub"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type StandaloneServer struct {
	grpcServer grpc.Server
	httpServer http.Server
	consumer   consumer.ConsumerServer
	mqConfig   configs.MQ
	logger     *zap.Logger
}

func NewStandaloneServer(grpcServer grpc.Server, httpServer http.Server, consumer consumer.ConsumerServer, mqConfig configs.MQ, logger *zap.Logger) *StandaloneServer {
	return &StandaloneServer{
		grpcServer: grpcServer,
		httpServer: httpServer,
		consumer:   consumer,
		mqConfig:   mqConfig,
		logger:     logger,
	}
}

func (s StandaloneServer) CreateTopic(topicName string) error {
	ctx := context.Background()
	logger := utils.LoggerWithContext(ctx, s.logger).With(zap.String("topic_name", topicName))

	pubsubClient, err := pubsub.NewClient(ctx, s.mqConfig.ProjectID)
	if err != nil {
		return fmt.Errorf("failed to create pubsub client: %w", err)
	}
	topic := pubsubClient.Topic(topicName)

	// Check if the topic exists
	exists, err := topic.Exists(ctx)
	if err != nil {
		logger.With(zap.Error(err)).Error("Failed to check topic existence")
		return status.Error(codes.Internal, "failed to check topic existence")
	}

	// Create the topic if it does not exist
	if !exists {
		logger.Warn("Topic does not exist, creating it...")
		topic, err = pubsubClient.CreateTopic(ctx, topicName)
		if err != nil {
			logger.With(zap.Error(err)).Error("Failed to create topic")
			return status.Error(codes.Internal, "failed to create topic")
		}
		logger.Info("Topic created successfully", zap.String("topic_name", topic.ID()))
	}
	return nil
}

func (s StandaloneServer) CreateSubscription(topicName string, subscriptionName string) error {
	ctx := context.Background()
	logger := utils.LoggerWithContext(ctx, s.logger).
		With(zap.String("topic_name", topicName)).
		With(zap.String("subscription_name", subscriptionName))

	pubsubClient, err := pubsub.NewClient(ctx, s.mqConfig.ProjectID)
	if err != nil {
		return fmt.Errorf("failed to create pubsub client: %w", err)
	}

	topic := pubsubClient.Topic(topicName)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	topicExists, err := topic.Exists(ctx)
	if err != nil {
		logger.With(zap.Error(err)).Error("Failed to check topic existence")
		return err
	}

	if !topicExists {
		logger.Error("Topic does not exist", zap.String("topic_name", topicName))
		return err
	}
	// Ensure the subscription exists
	subscription := pubsubClient.Subscription(subscriptionName)
	subExists, err := subscription.Exists(ctx)
	if err != nil {
		logger.With(zap.Error(err)).Error("Failed to check subscription existence")
		return err
	}

	if !subExists {
		logger.Warn("Subscription does not exist, creating it...", zap.String("subscription_name", subscriptionName))
		_, err = pubsubClient.CreateSubscription(ctx, subscriptionName, pubsub.SubscriptionConfig{
			Topic:       topic,
			AckDeadline: 10 * time.Second,
		})
		if err != nil {
			logger.With(zap.Error(err)).Error("Failed to create subscription")
			return err
		}
		logger.Info("Subscription created successfully", zap.String("subscription_name", subscriptionName))
	}
	return nil
}

func (s StandaloneServer) Start() error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := s.CreateTopic(producer.AISAccountTopic)
	if err != nil {
		return err
	}
	err = s.CreateSubscription(producer.AISAccountTopic, consumer.AISAccountSubscription)
	if err != nil {
		return err
	}

	go func() {
		grpcStartErr := s.grpcServer.Start(ctx)
		s.logger.With(zap.Error(grpcStartErr)).Info("grpc server stopped")
	}()

	go func() {
		httpStartErr := s.httpServer.Start(ctx)
		s.logger.With(zap.Error(httpStartErr)).Info("http server stopped")
	}()

	go func() {
		consumerStartErr := s.consumer.Start(ctx)
		s.logger.With(zap.Error(consumerStartErr)).Info("message queue consumer stopped")

	}()

	// Listen for OS termination signals (Ctrl+C, SIGTERM)
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// Block until a termination signal is received
	sig := <-sigChan
	fmt.Printf("\nReceived signal: %v. Shutting down...\n", sig)

	return nil
}
