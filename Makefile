DB_DRIVER=postgres
DB_SOURCE="postgresql://postgres:admin@localhost:5432/relibcadb?sslmode=disable&search_path=public"
MIGRATIONS_DIR=/mnt/d/Engineer/Server-adv/Project/ReLibca/src/infrastructure/migrations

migrate-init:
	migrate create -dir ${MIGRATIONS_DIR} -ext sql $(name)

migrate-up:
	migrate -path ${MIGRATIONS_DIR} -database ${DB_SOURCE} -verbose up

migrate-down:
	migrate -path ${MIGRATIONS_DIR} -database ${DB_SOURCE} -verbose down

migrate-fix:
	migrate -path ${MIGRATIONS_DIR} -database ${DB_SOURCE} force 0