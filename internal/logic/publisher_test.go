package logic_test

import (
	"ais_service/internal/dataaccess/mq/producer"
	"ais_service/internal/logic"
	"ais_service/mocks"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPublishAisAccount_Success(t *testing.T) {
	mockProducer := new(mocks.AccountProducer)
	publisher := logic.NewPublisher(mockProducer)
	ctx := context.Background()

	params := logic.PublishAisAccountParams{
		Account_id:     1,
		Account_name:   "Test Account",
		Account_type:   2,
		Account_status: 1,
	}

	expectedEvent := producer.AccountEvent{
		Account_id:     params.Account_id,
		Account_name:   params.Account_name,
		Account_type:   params.Account_type,
		Account_status: params.Account_status,
	}

	mockProducer.On("Produce", ctx, expectedEvent).Return(nil)

	output, err := publisher.PublishAisAccount(ctx, params)

	assert.NoError(t, err)
	assert.Equal(t, logic.PublishAisAccountOutput{}, output)
	mockProducer.AssertExpectations(t)
}

func TestPublishAisAccount_Failure(t *testing.T) {
	mockProducer := new(mocks.AccountProducer)
	publisher := logic.NewPublisher(mockProducer)
	ctx := context.Background()

	params := logic.PublishAisAccountParams{
		Account_id:     2,
		Account_name:   "Error Account",
		Account_type:   1,
		Account_status: 0,
	}

	expectedEvent := producer.AccountEvent{
		Account_id:     params.Account_id,
		Account_name:   params.Account_name,
		Account_type:   params.Account_type,
		Account_status: params.Account_status,
	}

	mockProducer.On("Produce", ctx, expectedEvent).Return(assert.AnError)

	output, err := publisher.PublishAisAccount(ctx, params)

	assert.Error(t, err)
	assert.Equal(t, logic.PublishAisAccountOutput{}, output)
	mockProducer.AssertExpectations(t)
}
