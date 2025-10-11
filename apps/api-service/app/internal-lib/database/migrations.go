package database

import (
	"context"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/rs/zerolog"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/migrate"
)

// RunMigrationsFromPath runs migrations from a directory path
func RunMigrationsFromPath(ctx context.Context, db *bun.DB, path string, logger *zerolog.Logger) error {
	migrations := migrate.NewMigrations()

	// Discover migrations from the specified directory
	if err := discoverMigrationsFromDir(migrations, path); err != nil {
		logger.Error().Err(err).Msgf("Failed to discover migrations from path: %s", path)
		return fmt.Errorf("failed to discover migrations from path %s: %w", path, err)
	}

	return RunMigrations(ctx, db, migrations, logger)
}

// RunMigrations runs database migrations
func RunMigrations(ctx context.Context, db *bun.DB, migrations *migrate.Migrations, logger *zerolog.Logger) error {
	migrator := migrate.NewMigrator(db, migrations)

	if err := migrator.Init(ctx); err != nil {
		logger.Error().Err(err).Msg("Failed to initialize migrator")
		return fmt.Errorf("failed to initialize migrator: %w", err)
	}

	group, err := migrator.Migrate(ctx)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to run migrations")
		return fmt.Errorf("failed to run migrations: %w", err)
	}

	if group.IsZero() {
		logger.Info().Msg("No new migrations to run")
	} else {
		logger.Info().Msgf("Migrated to %s", group)
	}

	return nil
}

func discoverMigrationsFromDir(migrations *migrate.Migrations, dirPath string) error {
	fsys := os.DirFS(dirPath)

	return fs.WalkDir(fsys, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() {
			return nil
		}

		// Only process .sql files
		if !strings.HasSuffix(path, ".sql") {
			return nil
		}

		fullPath := filepath.Join(dirPath, path)
		content, err := os.ReadFile(fullPath)
		if err != nil {
			return fmt.Errorf("failed to read migration file %s: %w", fullPath, err)
		}

		// Determine migration type from filename
		var migrationFunc migrate.MigrationFunc
		if strings.Contains(path, ".up.sql") {
			migrationFunc = func(ctx context.Context, db *bun.DB) error {
				_, err := db.ExecContext(ctx, string(content))
				return err
			}
		} else if strings.Contains(path, ".down.sql") {
			// Skip down migrations for now, or handle separately
			return nil
		} else {
			return nil
		}

		migrations.Add(migrate.Migration{
			Name: strings.TrimSuffix(filepath.Base(path), ".up.sql"),
			Up:   migrationFunc,
		})

		return nil
	})
}
