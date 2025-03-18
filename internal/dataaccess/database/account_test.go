package database

import (
	"ais_service/internal/configs"
	"ais_service/internal/generated/grpc/ais_api"
	"context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

var (
	testDBConfig = configs.Database{
		Username: "admin",
		Password: "admin",
		Host:     "localhost",
		Port:     5432,
		Database: "ais",
	}
)

func setupTestDB(t *testing.T) (*sql.DB, func()) {
	logger, _ := zap.NewDevelopment()
	db, cleanup, err := InitializeAndMigrateUpDB(testDBConfig, logger)
	require.NoError(t, err, "failed to set up test database")

	return db, cleanup
}

func TestAisAccountDataAccessor(t *testing.T) {
	db, cleanup := setupTestDB(t)
	defer cleanup() // Ensure cleanup happens after tests

	database := InitializeGoquDB(db)
	logger, _ := zap.NewDevelopment()
	accessor := NewAisAccountDataAccessor(database, logger)

	t.Run("CreateAisAccount", func(t *testing.T) {
		ctx := context.Background()
		account := AisAccount{
			Account_id:     1,
			Account_name:   "Test Account",
			Account_type:   ais_api.AccountType_CASA,
			Account_status: ais_api.Status_active,
		}

		id, err := accessor.CreateAisAccount(ctx, account)
		require.NoError(t, err)
		require.Equal(t, account.Account_id, id)
	})

	t.Run("GetAisAccountByID", func(t *testing.T) {
		ctx := context.Background()
		account, err := accessor.GetAisAccountByID(ctx, 1)
		require.NoError(t, err)
		require.Equal(t, "Test Account", account.Account_name)
	})

	t.Run("UpdateAisAccount", func(t *testing.T) {
		ctx := context.Background()
		updatedAccount := AisAccount{
			Account_id:     1,
			Account_name:   "Updated Name",
			Account_type:   ais_api.AccountType_CASA,
			Account_status: ais_api.Status_active,
		}

		err := accessor.UpdateAisAccount(ctx, updatedAccount)
		require.NoError(t, err)

		account, err := accessor.GetAisAccountByID(ctx, 1)
		require.NoError(t, err)
		require.Equal(t, "Updated Name", account.Account_name)
	})
}
