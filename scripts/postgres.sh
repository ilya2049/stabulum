#!/bin/bash

# Migriations
export MIGRATIONS_DIR=internal/infrastructure/postgres/migrations
export POSTGRES_DSN="host=127.0.0.1 port=5432 user=store-app password=password dbname=store-app-db sslmode=disable"

## Add a new migration in a sql file.
goose -dir=${MIGRATIONS_DIR} create init_schema sql

## Migrate the DB to the most recent version available.
goose -dir=${MIGRATIONS_DIR} postgres "${POSTGRES_DSN}" up

## Rollback the last migration.
goose -dir=${MIGRATIONS_DIR} postgres "${POSTGRES_DSN}" down