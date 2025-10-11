# Database Migrations with Bun ORM

This project uses **Bun ORM's built-in migration system** for managing database schema changes.

## 📁 Project Structure

```
apps/api-service/
├── cmd/
│   └── migrate/
│       └── main.go          # Migration CLI tool
├── migrations/
│   ├── migrations.go        # Migration registry
│   ├── 20241009000001_create_jobs_table.up.sql
│   └── 20241009000001_create_jobs_table.down.sql
```

## 🚀 Quick Start

### 1. Initialize Migration Tables

First time only - creates the `bun_migrations` and `bun_migration_locks` tables:

```bash
make migrate-init
```

### 2. Run Migrations

Apply all pending migrations:

```bash
make migrate-up
```

### 3. Check Migration Status

See which migrations have been applied:

```bash
make migrate-status
```

Output example:
```
Migration Status:
================
20241009000001_create_jobs_table - applied at 2024-10-09 14:23:45
20241010120000_add_user_table - pending
```

### 4. Rollback Migration

Rollback the last migration group:

```bash
make migrate-down
```

### 5. Create New Migration

Create a new migration file:

```bash
make migrate-create
# Then enter migration name when prompted, e.g., "add_user_table"
```

## 📝 Migration Format

Bun supports two migration formats:

### Option 1: Go-based Migrations (Recommended for complex logic)

```go
// In migrations/migrations.go
func init() {
    Migrations.MustRegister(func(ctx context.Context, db *bun.DB) error {
        // UP migration
        _, err := db.ExecContext(ctx, `
            CREATE TABLE users (
                id BIGSERIAL PRIMARY KEY,
                email VARCHAR(255) UNIQUE NOT NULL
            );
        `)
        return err
    }, func(ctx context.Context, db *bun.DB) error {
        // DOWN migration
        _, err := db.ExecContext(ctx, `DROP TABLE IF EXISTS users;`)
        return err
    })
}
```

### Option 2: SQL Files (Recommended for simple DDL)

Your existing SQL files work perfectly! They're already set up:
- `20241009000001_create_jobs_table.up.sql` - Creates tables
- `20241009000001_create_jobs_table.down.sql` - Drops tables

## 🔧 Configuration

The migration tool uses the same environment variables as your app:

```bash
# Job Scheduler Database Configuration
JOB_SCHEDULER_POSTGRES_DB_HOST=localhost
JOB_SCHEDULER_POSTGRES_DB_PORT=5432
JOB_SCHEDULER_POSTGRES_DB_USERNAME=postgres
JOB_SCHEDULER_POSTGRES_DB_PASSWORD=postgres
JOB_SCHEDULER_POSTGRES_DB_DATABASE=job_scheduler
JOB_SCHEDULER_POSTGRES_DB_SSLMODE=disable

# Or use connection URL
JOB_SCHEDULER_POSTGRES_DB_URL=postgres://user:pass@localhost:5432/job_scheduler?sslmode=disable
```

## 📋 Available Make Commands

```bash
make migrate-init      # Initialize migration tables (first time only)
make migrate-up        # Run all pending migrations
make migrate-down      # Rollback last migration
make migrate-status    # Show migration status
make migrate-create    # Create new migration file
```

## 🔐 Migration Locks

Bun automatically handles migration locks to prevent concurrent migrations. The `bun_migration_locks` table ensures only one migration runs at a time.

## 🎯 Best Practices

1. **Always test migrations** in development before production
2. **Write reversible migrations** - every UP should have a DOWN
3. **Use transactions** for data migrations
4. **Keep migrations small** - one logical change per migration
5. **Never modify existing migrations** that have been applied to production
6. **Use meaningful names** for migrations (e.g., `add_user_email_index`)

## 🔄 Migration Workflow

```bash
# 1. Create a new migration
make migrate-create
# Enter: add_user_email_index

# 2. Edit the generated migration file in migrations/migrations.go

# 3. Check what will be applied
make migrate-status

# 4. Apply the migration
make migrate-up

# 5. Verify it worked
make migrate-status

# 6. If something went wrong, rollback
make migrate-down
```

## 🐛 Troubleshooting

### "Migration tables not initialized"
Run: `make migrate-init`

### "No migrations to run"
Check that your migrations are registered in `migrations/migrations.go`

### "Failed to acquire migration lock"
Another migration is running, or a previous migration crashed. Check `bun_migration_locks` table.

### "Connection refused"
Verify your database is running and environment variables are correct.

## 📚 Advanced Usage

### Run migrations programmatically in your app

```go
import (
    "github.com/sdivyansh59/digantara-backend-golang-assignment/migrations"
    "github.com/uptrace/bun/migrate"
)

func runMigrations(db *bun.DB) error {
    migrator := migrate.NewMigrator(db, migrations.Migrations)
    
    if err := migrator.Init(ctx); err != nil {
        return err
    }
    
    if _, err := migrator.Migrate(ctx); err != nil {
        return err
    }
    
    return nil
}
```

### Add to your app startup (optional)

You can automatically run migrations when your app starts by adding this to your `ProvideJobSchedulerDB` function:

```go
// Auto-migrate on startup (optional - not recommended for production)
migrator := migrate.NewMigrator(db, migrations.Migrations)
if _, err := migrator.Migrate(ctx); err != nil {
    logger.Warn().Err(err).Msg("Failed to auto-migrate")
}
```

## 🎉 Benefits of Bun Migrations

✅ **Built-in** - No external tools needed
✅ **Version control** - Migrations are part of your codebase
✅ **Transactional** - Each migration runs in a transaction
✅ **Lock support** - Prevents concurrent migrations
✅ **Rollback support** - Easy to undo changes
✅ **Go-based** - Write complex migrations in Go
✅ **SQL support** - Or use simple SQL files
✅ **Status tracking** - Know what's applied and what's pending

