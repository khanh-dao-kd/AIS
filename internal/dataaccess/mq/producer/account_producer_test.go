package producer_test

import (
	"ais_service/internal/dataaccess/mq/producer"
	"ais_service/internal/generated/grpc/ais_api"
	"ais_service/mocks"
	"context"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestProduce_Success(t *testing.T) {
	mockClient := new(mocks.Client)
	producerInstance := producer.NewAccountProducer(mockClient, &zap.Logger{})
	ctx := context.Background()

	event := producer.AccountEvent{
		Account_id:     1,
		Account_name:   "Test Account",
		Account_type:   ais_api.AccountType_CASA,
		Account_status: ais_api.Status_active,
	}

	eventBytes, _ := json.Marshal(event)

	mockClient.On("Produce", ctx, producer.AISAccountTopic, eventBytes).Return(nil)

	err := producerInstance.Produce(ctx, event)

	assert.NoError(t, err)
	mockClient.AssertExpectations(t)
}

// func TestProduce_MarshalError(t *testing.T) {
// 	mockClient := new(mocks.Client)
// 	producerInstance := producer.NewAccountProducer(mockClient)
// 	ctx := context.Background()

// 	event := producer.AccountEvent{
// 		Account_id:     1,
// 		Account_name:   "Test Account",
// 		Account_type:   ais_api.AccountType(-1),
// 		Account_status: ais_api.Status(-1),
// 	}

// 	err := producerInstance.Produce(ctx, event)

// 	assert.Error(t, err)
// 	assert.Equal(t, codes.Internal, status.Code(err))
// 	assert.Contains(t, err.Error(), "failed to marshal download task created event")

// 	mockClient.AssertNotCalled(t, "Produce", mock.Anything, mock.Anything, mock.Anything)
// }

func TestProduce_ClientError(t *testing.T) {
	mockClient := new(mocks.Client)
	producerInstance := producer.NewAccountProducer(mockClient, &zap.Logger{})
	ctx := context.Background()

	event := producer.AccountEvent{
		Account_id:     2,
		Account_name:   "Error Account",
		Account_type:   ais_api.AccountType_CASA,
		Account_status: ais_api.Status_active,
	}

	eventBytes, _ := json.Marshal(event)

	mockClient.On("Produce", ctx, producer.AISAccountTopic, eventBytes).Return(status.Error(codes.Internal, "client error"))

	err := producerInstance.Produce(ctx, event)

	assert.Error(t, err)
	assert.Equal(t, codes.Internal, status.Code(err))
	assert.Contains(t, err.Error(), "failed to marshal download task created event")

	mockClient.AssertExpectations(t)
}
