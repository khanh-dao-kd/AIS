package database

import (
	"ais_service/internal/generated/grpc/ais_api"
	"context"

	"github.com/doug-martin/goqu/v9"
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
	Account_id     uint64
	Account_name   string
	Account_type   ais_api.AccountType
	Account_status ais_api.Status
	Source         string
}

type AisAccountDataAccessor interface {
	CreateAisAccount(ctx context.Context, account AisAccount) (uint64, error)
	GetAisAccountByID(ctx context.Context, account_id uint64) (AisAccount, error)
	UpdateAisAccount(ctx context.Context, account AisAccount) error
}

type aisAccountDataAccessor struct {
	database Database
}

func NewAisAccountDataAccessor(database Database) AisAccountDataAccessor {
	return &aisAccountDataAccessor{
		database: database,
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
		return 0, status.Error(codes.Internal, "failed to create ais account")
	}
	return account.Account_id, nil
}

func (a aisAccountDataAccessor) GetAisAccountByID(ctx context.Context, account_id uint64) (AisAccount, error) {
	account := AisAccount{}

	found, err := a.database.
		From(TabAisAccount).
		Where(goqu.C(ColNameAccountID).Eq(account.Account_id)).
		ScanStructContext(ctx, account)
	if err != nil {
		return AisAccount{}, status.Error(codes.Internal, "failed to get ais account by id")
	}
	if !found {
		return AisAccount{}, status.Error(codes.NotFound, "account not found")
	}
	return account, nil
}

func (a aisAccountDataAccessor) UpdateAisAccount(ctx context.Context, account AisAccount) error {
	_, err := a.database.
		Update(TabAisAccount).
		Set(account).
		Where(goqu.C(ColNameAccountID).Eq(account.Account_id)).
		Executor().
		ExecContext(ctx)
	if err != nil {
		return status.Error(codes.Internal, "failed to update ais account")
	}
	return nil
}
