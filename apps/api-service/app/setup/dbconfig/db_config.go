package dbconfig

import (
	"context"
	"fmt"
	"time"

	"github.com/rs/zerolog"
	"github.com/sdivyansh59/digantara-backend-golang-assignment/app/internal-lib/database"
	"github.com/sdivyansh59/digantara-backend-golang-assignment/app/internal-lib/utils"
	"github.com/uptrace/bun"
)

type JobSchedulerDB struct {
	*bun.DB
}

func ProvideJobSchedulerDB(logger *zerolog.Logger, withLogger *utils.WithLogger, config *utils.DefaultConfig) (*JobSchedulerDB, error) {
	jobSchedulerDBConfig := *config
	jobSchedulerDBConfig.ServicePrefix = "JOB_SCHEDULER"

	db := database.NewPostgres(&jobSchedulerDBConfig, logger)

	if utils.IsDebug() {
		db.AddQueryHook(database.NewDBLogger(withLogger))
	}

	// Verify database connection health
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		logger.Error().Err(err).Msg("Database health check failed")
		return nil, fmt.Errorf("database health check failed: %w", err)
	}

	logger.Info().Msg("Job Scheduler Database health check passed")

	return &JobSchedulerDB{DB: db}, nil
}

func AddJobSchedulerDBMigrationsHook(db *JobSchedulerDB, logger *zerolog.Logger) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Get migrations path from environment or use default
	migrationsPath := utils.GetEnvOr("MIGRATIONS_PATH", "migrations")

	logger.Info().Msgf("Running migrations from path: %s", migrationsPath)

	// Run migrations
	if err := database.RunMigrationsFromPath(ctx, db.DB, migrationsPath, logger); err != nil {
		logger.Error().Err(err).Msg("Failed to run migrations")
		return fmt.Errorf("failed to run migrations: %w", err)
	}

	logger.Info().Msg("Migrations completed successfully")
	return nil
}
