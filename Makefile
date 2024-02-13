include .env

new_migration:
	migrate create -ext sql -dir db/migration -seq $(name)

migrate_up:
	migrate -path db/migration -database ${PSQL_MIGRATION_URL} -verbose up

migrate_up_1:
	migrate -path db/migration -database ${PSQL_MIGRATION_URL} -verbose up 1

migrate_down:
	migrate -path db/migration -database ${PSQL_MIGRATION_URL} -verbose down

migrate_down_1:
	migrate -path db/migration -database ${PSQL_MIGRATION_URL} -verbose down 1

db_docs:
	dbdocs build db/database.dbml