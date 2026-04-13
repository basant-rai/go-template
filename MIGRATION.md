# Database Migration System

## Overview

The project uses **golang-migrate** — a pure Go migration tool with CLI support.

- **Manual control** — You decide when to apply migrations (not automatic)
- **CLI-based** — Use `migrate` command line tool
- **Atomic transactions** — Each migration succeeds or fails completely
- **Idempotent** — Safe to re-run (uses `IF NOT EXISTS` and `IF EXISTS`)
- **Rollback support** — Every migration has up and down scripts
- **Tracks state** — Applied migrations recorded in `schema_migrations` table

## Workflow

1. **Developer creates** → `make migrate-create NAME=my_feature`
2. **Edit migration files** → Write SQL in `.up.sql` and `.down.sql`
3. **Apply migrations** → `make migrate-up`
4. **Check status** → `make migrate-status`
5. **Rollback if needed** → `make migrate-down`

## Quick Start

### 1. Install golang-migrate CLI

```bash
make migrate-install

make get https://pkg.go.dev/github.com/basant-rai/gomigrate
```

### 2. Apply Migrations

```bash
make migrate-install
```

## Using Make Commands

All migration commands are available via Makefile:

```bash
# Install migrate CLI
make migrate-install

# Initialize migration
make migrate-init

# Create a new migration
make migrate-generate NAME={{migration_name}};
#make migrate-create NAME=add_new_table

# Check new migration
make migrate-diff

# Auto generate migration
make migrate-generate NAME={{migration_name}};

# Apply all pending migrations
make migrate-up

# Rollback one migration
make migrate-down

# Show current migration version
make migrate-status

# Reset database (dev only)
make migrate-reset

# Build and run server
make dev
```

## Current Migrations

Located in `migrations/` directory. Format: `NNNNNN_name_{up,down}.sql`

```
✅ 000001_initial_schema.{up,down}.sql
   - Core tables: users, other tables...
```

## Best Practices

### DO ✅

- Use `IF NOT EXISTS` and `IF EXISTS` for idempotency
- Create indexes for foreign keys and frequently queried columns
- Include CHECK constraints for enum-like fields
- Write both `.up.sql` AND `.down.sql` migrations
- Test rollbacks in development
- Include a status message (final SELECT) in each migration
- Use descriptive migration names (snake_case)
- Add comments explaining complex changes

### DON'T ❌

- Never edit an already-applied migration
- Don't forget the down migration
- Don't use raw SQL without `IF NOT EXISTS`
- Don't make destructive changes lightly
- Don't store raw secrets (we encrypt sensitive fields)
- Don't run migrations during peak traffic

## Troubleshooting

### Check Applied Migrations

```sql
SELECT * FROM schema_migrations ORDER BY version;
```

### Migration Failed (Dirty State)

If a migration fails, the database is marked "dirty" and requires manual intervention:

```sql
-- Check dirty status
SELECT * FROM schema_migrations WHERE dirty = true;

-- Reset dirty flag (dev only)
UPDATE schema_migrations SET dirty = false WHERE version = 5;

-- Or drop and retry
DELETE FROM schema_migrations WHERE version = 5;
```

Then fix the migration SQL and try again.

### Full Reset (Development Only)

```bash
# Delete all migrations
migrate -path migrations -database "$DATABASE_URL" down -all

# Or manually
psql -c "DROP TABLE schema_migrations;" $DATABASE_URL
```

## Production Checklist

Before running migrations in production:

- [ ] Test migration in staging environment first
- [ ] Backup database
- [ ] Schedule downtime if needed
- [ ] Review migration SQL for performance issues
- [ ] Have rollback plan ready
- [ ] Monitor database during migration
- [ ] Verify data integrity after migration

## Environment Variables

Set these before running migrations:

```bash
export DATABASE_URL="postgres://user:password@localhost:5432/db?sslmode=disable"
```

Or add to `.env`:

```
DATABASE_URL=postgres://user:password@localhost:5432/db?sslmode=disable
```

## References

- [golang-migrate Documentation](https://github.com/golang-migrate/migrate)
- [PostgreSQL Documentation](https://www.postgresql.org/docs/)
