include .env
export $(shell sed 's/=.*//' .env)

build:
	go build -C cmd/valkyrie-db/ -o ../../bin/valkyrie-db

test:
	go test -v ./...

run: build
	./bin/valkyrie-db
