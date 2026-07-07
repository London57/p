sqlc-generate:
	@docker compose run --rm sqlc generate

init-postgres-env:
	@mkdir -p "out/pgdata" && \
	sudo chown -R 999:999 "out/pgdata" &&\
	sudo chmod -R 700 "out/pgdata";
	
pg-up:
	@docker compose up postgres 