ENV_FILE := $(if $(wildcard .env),.env,.example.env)

include $(ENV_FILE)
export $(shell sed 's/=.*//' $(ENV_FILE))

build:
	go build -C cmd/valkyrie-db/ -o ../../bin/valkyrie-db

test:
	go test -v ./...

run: build
	./bin/valkyrie-db
