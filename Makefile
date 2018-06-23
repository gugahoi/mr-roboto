REPOSITORY := gugahoi/mr-roboto
TAG ?= latest

build:
	docker build -t $(REPOSITORY):$(TAG) .

run: build
	docker run $(REPOSITORY):$(TAG)
