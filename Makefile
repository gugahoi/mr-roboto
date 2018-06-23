REPOSITORY := gugahoi/mr-roboto
TAG ?= latest

build:
	docker build -t $(REPOSITORY):$(TAG) .
