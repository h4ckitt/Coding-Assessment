SHELL := /bin/bash


test:
	source ./db_test.env
	go test -cover ./domain
	go test -cover ./helpers
	go test ./infrastructure/db/postgres
	go test ./repository
	go test -cover ./usecases
	go test -cover ./adapter/grpc
	go test -cover ./adapter/rest
