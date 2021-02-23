.PHONY: build check clean image publish publish-latest test

TAG_NAME := $(shell git tag -l --contains HEAD)
SHA := $(shell git rev-parse --short HEAD)
VERSION := $(if $(TAG_NAME),$(TAG_NAME),$(SHA))
BUILD_DATE := $(shell date -u '+%Y-%m-%d_%I:%M:%S%p')
DOCKER_REGISTRY := gcr.io
DOCKER_REPOSITORY := trois-six/fauna-exporter

default: clean check test build

clean:
	rm -rf cover.out

test: clean
	go test -v -cover ./...

build: clean
	@echo Version: $(VERSION) $(BUILD_DATE)
	CGO_ENABLED=0 go build -v -ldflags '-X "main.version=${VERSION}" -X "main.commit=${SHA}" -X "main.date=${BUILD_DATE}"'

image:
	docker build -t $(DOCKER_REGISTRY)/$(DOCKER_REPOSITORY):$(VERSION) .

publish:
	docker push $(DOCKER_REGISTRY)/$(DOCKER_REPOSITORY):$(VERSION)

publish-latest:
	docker tag $(DOCKER_REGISTRY)/$(DOCKER_REPOSITORY):$(VERSION) $(DOCKER_REGISTRY)/$(DOCKER_REPOSITORY):latest
	docker push $(DOCKER_REGISTRY)/$(DOCKER_REPOSITORY):latest

check:
	golangci-lint run
