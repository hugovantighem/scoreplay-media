install-deps:
	go install github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@latest
	go install go.uber.org/mock/mockgen@latest

generate:
	go generate ./...

generate-scratch: install-deps generate

test-unit:
	go test ./app/... ./infra/...

test-unit-scratch: generate-scratch test-unit

test-integration:
	go test ./tests/...

build: 
	go build -o main .

build-scratch: generate-scratch build

run:
	./main

run-scratch: build-scratch run

test-unit-coverage: generate-scratch
	go test -coverpkg=./app/...,./infra/...  -coverprofile=profile.cov ./app/... ./infra/... 
	go tool cover -func profile.cov

test-all-coverage: generate-scratch
	go test -coverpkg=./app/...,./infra/...  -coverprofile=profile.cov ./...
	go tool cover -func profile.cov

copy-env-file:
	cp .env-example .env

start-db:
	docker run --name mongodb -p 27017:27017 -d mongodb/mongodb-community-server:latest
