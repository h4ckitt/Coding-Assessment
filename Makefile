SHELL := /bin/bash
include .env

.PHONY: run
run:
	@source .env
	@go run main.go

.PHONY: build
build:
	@GOOS=linux GOARCH=amd64 go build -o ./bin/assessment/assessment-linux
	@GOOS=darwin GOARCH=amd64 go build -o ./bin/assessment/assessment-darwin

.PHONY: test
test:
	@source .env
	@go test -cover ./domain
	@go test -cover ./helpers
	@go test ./infrastructure/db/postgres
	@go test ./repository
	@go test -cover ./usecases
	@go test -cover ./adapter/grpc
	@go test -cover ./adapter/rest

.PHONY: docker
docker: Dockerfile docker-compose.yml
	@docker-compose up

.PHONY: clean
clean:
	@docker-compose down --volumes
	@docker image rm -f area99_assessment:dharmy
