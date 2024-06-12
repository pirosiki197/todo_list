#!/bin/sh

set -eu

SQL_FILE=$1

DB_HOST=${DB_HOST:-localhost}
DB_PORT=${DB_PORT:-3306}
DB_USER=${DB_USER:-root}
DB_PASS=${DB_PASS:-root}
DB_NAME=${DB_NAME:-todo_list}

mysqldef --host="$DB_HOST" --port="$DB_PORT" --user="$DB_USER" --password="$DB_PASS" "$DB_NAME" < "$SQL_FILE"
