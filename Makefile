.PHONY: local docker test

-include .env

GOCMD=go
GOBUILD=$(GOCMD) build
GORUN=$(GOCMD) run
GOGENERATE=$(GOCMD) generate
NAME=correspondence-composer
ENTRY_PATH=cmd/$(NAME)/main.go

run:
	${GORUN} ${ENTRY_PATH}

build:
	$(GOBUILD) -o bin/$(NAME) $(ENTRY_PATH)

build-linux:
	 env GOOS=linux $(GOBUILD) -o bin/$(NAME) -i $(ENTRY_PATH)

kafka-start:
	docker-compose -f kafka.yml up -d --remove-orphans

kafka-stop:
	docker-compose -f kafka.yml down

docker-build:
	docker-compose -f docker-compose-local.yml build correspondence-composer

docker-run:
	docker-compose -f docker-compose-local.yml up

test:
	$(GOCMD) test -v ./... -p 1

lint:
	golangci-lint run -c .golangci.yml --fix

generate-xsd-types:
	xgen -i "./xsds/$(xsd).xsd" -o "./models/generated/$(output).go" -l Go

