SHELL := /bin/bash

run:
	source .env
	go run main.go

build:
	GOOS=linux GOARCH=amd64 go build -o ./bin/assessment/assessment-linux
	GOOS=darwin GOARCH=amd64 go build -o ./bin/assessment/assessment-darwin

test:
	source .env
	go test -cover ./domain
	go test -cover ./helpers
	go test ./infrastructure/db/postgres
	go test ./repository
	go test -cover ./usecases
	go test -cover ./adapter/grpc
	go test -cover ./adapter/rest

docker: Dockerfile docker-compose.yml
	docker-compose up

clean:
	docker-compose down --volumes
	docker image rm -f area99_web
