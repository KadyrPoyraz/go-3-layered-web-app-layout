run.app:
	go run cmd/main.go

run.db:
	docker-compose -f docker/docker-compse.yml up -d db

run.db.migrate:
	goose -dir internal/migrations -allow-missing postgres "host=localhost port=1337 user=user password=228 dbname=db sslmode=disable" up

stop.db:
	docker-compose -f docker/docker-compse.yml  db
