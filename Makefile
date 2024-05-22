.PHONY: build run test docker-up docker-down docker-restart

build:
	@ go build -o bin/server cmd/server/main.go

run: build
	@ ./bin/server

test:
	@ go test -count=1 ./...

docker-up:
	@ docker-compose up -d

docker-down:
	@ docker-compose down

docker-restart:
	@ docker-down docker-up
