REPOSITORY := gugahoi/mr-roboto
TAG ?= latest

all: test

test:
	go test -cover -v ./...

clean:
	rm -rf build/

build: clean
	go build -o build/mr-roboto ./...

docker-build:
	docker build -t $(REPOSITORY):$(TAG) .

docker-run: docker-build
	docker run $(REPOSITORY):$(TAG)
