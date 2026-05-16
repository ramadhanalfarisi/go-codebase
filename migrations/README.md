# Migrations Documentation

This folder contains SQL migrations for the project and uses `golang-migrate` to apply them.

## Purpose

The migration runner in this repository is implemented in `app/migrate/migrate.go`. It uses the `golang-migrate` library to apply all up migrations from the `migrations/` directory.

## Required Environment Variables

The application loads migration configuration from environment variables.

Required vars:

- `DB_URL` — PostgreSQL connection string used by the Go database driver.
- `DB_NAME` — database name used by the migrate library when opening the database instance.
- `MIGRATIONS_PATH` — file source path for migration files.

Example values from `.env.example`:

```env
DB_URL=postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable&search_path=codebase
DB_NAME=postgres
MIGRATIONS_PATH=file://./migrations/
```

## How It Works

`app/migrate/migrate.go` performs the following steps:

1. Connects to the database using `drivers.ConnectDB()` and `DB_URL`.
2. Creates a `postgres` migration driver instance from the open SQL connection.
3. Loads migration files from `MIGRATIONS_PATH`.
4. Runs `m.Up()` to execute all pending `*.up.sql` files.

If a migration fails, the program logs the error and exits.

## Running Migrations

From the repository root, set your environment variables and run:

```bash
go run main.go migrate
```

This will execute all pending migrations in the `migrations/` folder.

## Migration File Naming

Migration files must follow the `golang-migrate` naming convention:

- `0001_description.up.sql`
- `0001_description.down.sql`

Existing files in this repo:

- `1_create_user_table.up.sql`
- `1_create_user_table.down.sql`
- `2_create_product_table.up.sql`
- `2_create_product_table.down.sql`

## Using the `golang-migrate` CLI

If you prefer to manage migrations manually with the CLI, install it first:

```bash
curl -L https://github.com/golang-migrate/migrate/releases/latest/download/migrate.windows-amd64.tar.gz | tar xvz
```

Then run commands from the repository root:

```bash
migrate -path ./migrations -database "$DB_URL" up
```

To step down one migration:

```bash
migrate -path ./migrations -database "$DB_URL" down 1
```

To force a specific version:

```bash
migrate -path ./migrations -database "$DB_URL" force <version>
```

## Recommended Setup

Use the provided `.env.example` as a template and ensure the following values are set before running migrations:

- `DB_URL`
- `DB_NAME`
- `MIGRATIONS_PATH`

A typical value for `MIGRATIONS_PATH` is:

```env
MIGRATIONS_PATH=file://./migrations/
```

## Notes

- The app migration runner only supports applying `up` migrations via `m.Up()`.
- Rollback behavior is not exposed through the `go run main.go migrate` command in this app.
- The `golang-migrate` CLI can still be used for more advanced operations such as `down`, `force`, and `version`.
