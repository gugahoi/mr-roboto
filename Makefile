REPOSITORY := gugahoi/mr-roboto
TAG ?= latest

all: test

test:
	go test -v ./...

build:
	go build -o build/mr-roboto ./...

docker-build:
	docker build -t $(REPOSITORY):$(TAG) .

docker-run: docker-build
	docker run $(REPOSITORY):$(TAG)
