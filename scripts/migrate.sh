#!/bin/sh
set -e

echo "Running menu-service migrations..."

psql "$DATABASE_URL" <<'EOSQL'
CREATE TABLE IF NOT EXISTS menu_items (
	id SERIAL PRIMARY KEY,
	name TEXT NOT NULL,
	description TEXT,
	price NUMERIC(10,2) NOT NULL
);
EOSQL

echo "Menu migrations complete."
