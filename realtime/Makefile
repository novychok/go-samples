include .env

.PHONY: migrate-new
migrate-new:
	goose -dir db/migrations create ${name} sql

migrate-up:
	goose -dir ./db/migrations postgres "user=${PSQL_USER} dbname=${PSQL_DBNAME} password=${PSQL_PASS} \
	host=${PSQL_HOST} sslmode=${PSQL_SSLMODE}" up

migrate-down:
	goose -dir ./db/migrations postgres "user=${PSQL_USER} dbname=${PSQL_DBNAME} password=${PSQL_PASS} \
	host=${PSQL_HOST} sslmode=${PSQL_SSLMODE}" down
