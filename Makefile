all: build run

build:
	go build ./cmd/main

run:
	go run ./cmd/main

clean:
	go mod tidy

test:
	go test -v -count=1 ./...

test-cover:
	go test -v -count=1 ./... -covermode=count

test-out:
	go test -v -coverprofile cover.out ./...
	go tool cover -html=cover.out -o cover.html

generate-mocks:
	mockery --all --output internal/mocks

rebuild-env: stop-env destroy-env build-env

start-env:
	docker-compose -f ./docker-compose.yml up -d

stop-env:
	docker-compose -f ./docker-compose.yml stop

destroy-env:
	docker-compose -f ./docker-compose.yml rm -f

build-env:
	docker-compose -f ./docker-compose.yml build
