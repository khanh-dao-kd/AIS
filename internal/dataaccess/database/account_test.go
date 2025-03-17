package database_test

import (
	"ais_service/internal/dataaccess/database"
	"ais_service/mocks"
	"context"
	"testing"

	"github.com/doug-martin/goqu/v9"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.uber.org/zap"
)

func TestCreateAisAccount(t *testing.T) {
	mockDB := new(mocks.Database)
	accessor := database.NewAisAccountDataAccessor(mockDB, &zap.Logger{})

	account := database.AisAccount{
		Account_id:     1,
		Account_name:   "Test Account",
		Account_type:   1,
		Account_status: 2,
	}

	// Create a valid InsertDataset
	mockInsert := goqu.New("postgres", nil).Insert(database.TabAisAccount)
	mockDB.On("Insert", database.TabAisAccount).Return(mockInsert)

	// Mock ExecContext on the Executor
	mockDB.On("ExecContext", mock.Anything, mock.Anything).Return(nil, nil)

	id, err := accessor.CreateAisAccount(context.Background(), account)
	assert.NoError(t, err)
	assert.Equal(t, account.Account_id, id)
	mockDB.AssertExpectations(t)
}

func TestGetAisAccountByID(t *testing.T) {
	mockDB := new(mocks.Database)
	accessor := database.NewAisAccountDataAccessor(mockDB, &zap.Logger{})

	account := database.AisAccount{
		Account_id:     1,
		Account_name:   "Test Account",
		Account_type:   1,
		Account_status: 0,
	}

	mockSelect := new(goqu.SelectDataset)
	mockDB.On("From", database.TabAisAccount).Return(mockSelect)
	mockDB.On("ScanStructContext", mock.Anything, mock.Anything).Run(func(args mock.Arguments) {
		arg := args.Get(1).(*database.AisAccount)
		*arg = account
	}).Return(true, nil)

	retrievedAccount, err := accessor.GetAisAccountByID(context.Background(), account.Account_id)
	assert.NoError(t, err)
	assert.Equal(t, account, retrievedAccount)
	mockDB.AssertExpectations(t)
}

func TestUpdateAisAccount(t *testing.T) {
	mockDB := new(mocks.Database)
	logger := zap.NewNop()
	accessor := database.NewAisAccountDataAccessor(mockDB, logger)

	account := database.AisAccount{
		Account_id:     1,
		Account_name:   "Test Account",
		Account_type:   1,
		Account_status: 0,
	}

	mockUpdate := new(goqu.UpdateDataset)
	mockDB.On("Update", database.TabAisAccount).Return(mockUpdate)
	mockDB.On("ExecContext", mock.Anything).Return(nil, nil)

	err := accessor.UpdateAisAccount(context.Background(), account)
	assert.NoError(t, err)
	mockDB.AssertExpectations(t)
}
