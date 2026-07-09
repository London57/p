include .env
export

export PROJECT_ROOT=$(shell pwd)

sqlc-generate:
	@docker compose run --rm sqlc generate


init-postgres-env:
	@mkdir -p "out/pgdata" && \
	sudo chown -R 999:999 "out/pgdata" &&\
	sudo chmod -R 700 "out/pgdata";
	
pg-up:
	@docker compose up postgres 

migrate-create:
	@if [ -z "$(seq)" ]; then \
		echo "Отсутствует обязательный параметр 'seq'"; \
		exit 1; \
	fi

	@docker compose run --rm postgres-migrate \
		create \
		-ext sql \
		-dir /migrations \
		-seq "$(seq)";

migrate-action:
	@if [ -z "$(action)" ]; then \
		echo "Не задан параметр 'action'"; \
		exit 1; \
	fi
 
 
	@docker compose run --rm postgres-migrate \
		-path /migrations \
		-database postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@postgres:5432/${POSTGRES_DB}?sslmode=disable \
		"$(action)"