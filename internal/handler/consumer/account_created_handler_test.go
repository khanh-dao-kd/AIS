package consumer_test

import (
	"ais_service/internal/dataaccess/mq/producer"
	"ais_service/internal/handler/consumer"
	"ais_service/internal/logic"
	"ais_service/mocks"
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAccountCreatedHandler_Handle_Success(t *testing.T) {
	mockLogic := new(mocks.AccountLogic) // Mock AccountLogic
	handler := consumer.NewAccountCreatedHandler(mockLogic)
	ctx := context.Background()
	event := producer.AccountEvent{
		Account_id:     1,
		Account_name:   "Test Account",
		Account_type:   0,
		Account_status: 1,
	}

	mockLogic.On("CreateAisAccount", ctx, mock.Anything).
		Return(logic.CreateAisAccountOutput{}, nil).Once()

	err := handler.Handle(ctx, event)
	assert.NoError(t, err)

	mockLogic.AssertExpectations(t)
}

func TestAccountCreatedHandler_Handle_Failure(t *testing.T) {
	mockLogic := new(mocks.AccountLogic) // Mock AccountLogic
	handler := consumer.NewAccountCreatedHandler(mockLogic)
	ctx := context.Background()
	event := producer.AccountEvent{
		Account_id:     2,
		Account_name:   "Failing Account",
		Account_type:   1,
		Account_status: 0,
	}

	mockLogic.On("CreateAisAccount", ctx, mock.Anything).
		Return(0, errors.New("failed to create account")).Once()

	err := handler.Handle(ctx, event)
	assert.Error(t, err)
	assert.Equal(t, "failed to create account", err.Error())

	mockLogic.AssertExpectations(t)
}
