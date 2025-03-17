package database

import (
	"ais_service/internal/generated/grpc/ais_api"
	"context"

	"github.com/doug-martin/goqu/v9"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	TabAisAccount = goqu.T("ais_account")
)

const (
	ColNameAccountID     = "account_id"
	ColNameAccountName   = "account_name"
	ColNameAccountType   = "account_type"
	ColNameAccountStatus = "account_status"
)

type AisAccount struct {
	Account_id     uint64              `db:"account_id"`
	Account_name   string              `db:"account_name"`
	Account_type   ais_api.AccountType `db:"account_type"`
	Account_status ais_api.Status      `db:"account_status"`
}

type AisAccountDataAccessor interface {
	CreateAisAccount(ctx context.Context, account AisAccount) (uint64, error)
	GetAisAccountByID(ctx context.Context, account_id uint64) (AisAccount, error)
	GetAisAccountByIDForUpdate(ctx context.Context, account_id uint64) (AisAccount, error)
	UpdateAisAccount(ctx context.Context, account AisAccount) error
}

type aisAccountDataAccessor struct {
	database Database
	logger   *zap.Logger
}

func NewAisAccountDataAccessor(database Database, logger *zap.Logger) AisAccountDataAccessor {
	return &aisAccountDataAccessor{
		database: database,
		logger:   logger,
	}
}

func (a aisAccountDataAccessor) CreateAisAccount(ctx context.Context, account AisAccount) (uint64, error) {
	_, err := a.database.
		Insert(TabAisAccount).
		Rows(goqu.Record{
			ColNameAccountID:     account.Account_id,
			ColNameAccountName:   account.Account_name,
			ColNameAccountType:   account.Account_type,
			ColNameAccountStatus: account.Account_status,
		}).
		Executor().ExecContext(ctx)
	if err != nil {
		a.logger.With(zap.Error(err)).Error("failed to create account")
		return 0, status.Error(codes.Internal, "failed to create ais account")
	}
	return account.Account_id, nil
}

func (a aisAccountDataAccessor) GetAisAccountByID(ctx context.Context, account_id uint64) (AisAccount, error) {
	account := AisAccount{}

	found, err := a.database.
		From(TabAisAccount).
		Where(goqu.C(ColNameAccountID).Eq(account_id)).
		ScanStructContext(ctx, &account)
	if err != nil {
		a.logger.With(zap.Error(err)).Error("failed to get last inserted id")
		return AisAccount{}, status.Error(codes.Internal, "failed to get ais account by id")
	}
	if !found {
		a.logger.With(zap.Error(err)).Error("account not found")
		return AisAccount{}, status.Error(codes.NotFound, "account not found")
	}
	return account, nil
}

func (a aisAccountDataAccessor) GetAisAccountByIDForUpdate(ctx context.Context, account_id uint64) (AisAccount, error) {
	account := AisAccount{}

	found, err := a.database.
		From(TabAisAccount).
		Where(goqu.C(ColNameAccountID).Eq(account_id)).
		ForUpdate(goqu.Wait).
		ScanStructContext(ctx, &account)
	if err != nil {
		a.logger.With(zap.Error(err)).Error("failed to get last inserted id")
		return AisAccount{}, status.Error(codes.Internal, "failed to get ais account by id")
	}
	if !found {
		a.logger.With(zap.Error(err)).Error("account not found")
		return AisAccount{}, status.Error(codes.NotFound, "account not found")
	}
	return account, nil
}

func (a aisAccountDataAccessor) UpdateAisAccount(ctx context.Context, account AisAccount) error {
	_, err := a.database.
		Update(TabAisAccount).
		Set(goqu.Record{
			"account_name":   account.Account_name,
			"account_type":   account.Account_type,
			"account_status": account.Account_status,
		}).
		Where(goqu.C(ColNameAccountID).Eq(account.Account_id)).
		Executor().
		ExecContext(ctx)
	if err != nil {
		a.logger.With(zap.Error(err)).Error("failed to update ais account")
		return status.Error(codes.Internal, "failed to update ais account")
	}
	return nil
}
