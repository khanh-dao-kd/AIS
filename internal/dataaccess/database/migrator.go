package database

import (
	"ais_service/internal/utils"
	"context"
	"database/sql"
	"embed"

	migrate "github.com/rubenv/sql-migrate"
	"go.uber.org/zap"
)

var (
	//go:embed migrations/postgres/*
	migrationDirectory embed.FS
)

type Migrator interface {
	Up(ctx context.Context) error
	Down(ctx context.Context) error
}
type migrator struct {
	db     *sql.DB
	logger *zap.Logger
}

func NewMigrator(db *sql.DB, logger *zap.Logger) Migrator {
	return &migrator{
		db:     db,
		logger: logger,
	}
}
func (m migrator) migrate(ctx context.Context, direction migrate.MigrationDirection) error {
	logger := utils.LoggerWithContext(ctx, m.logger).With(zap.Int("direction", int(direction)))

	migrationCount, err := migrate.ExecContext(ctx, m.db, "postgres", migrate.EmbedFileSystemMigrationSource{
		FileSystem: migrationDirectory,
		Root:       "migrations/postgres",
	}, direction)
	if err != nil {
		logger.With(zap.Error(err)).Error("failed to execute migration")
		return err
	}
	logger.With(zap.Int("migration_count", migrationCount)).Info("successfully executed database migrations")

	return nil
}
func (m migrator) Down(ctx context.Context) error {
	return m.migrate(ctx, migrate.Down)
}
func (m migrator) Up(ctx context.Context) error {
	return m.migrate(ctx, migrate.Up)
}
