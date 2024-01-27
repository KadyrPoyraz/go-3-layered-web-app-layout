run-postgresql:
	docker-compose -f docker/docker-compse.yml up -d db

run:
	go run cmd/main.go

migrate:
	goose -dir internal/migrations -allow-missing postgres "host=localhost port=1337 user=user password=228 dbname=db sslmode=disable" up
