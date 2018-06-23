REPOSITORY := gugahoi/mr-roboto
TAG ?= latest

all: test

test:
	go test -v ./...

build:
	docker build -t $(REPOSITORY):$(TAG) .

run: build
	docker run $(REPOSITORY):$(TAG)
