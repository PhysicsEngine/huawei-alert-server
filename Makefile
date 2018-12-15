GIT_REF := $(shell git describe --always --tag)
VERSION ?= commit-$(GIT_REF)

SERVICE_NAME := $(shell basename ${PWD})

GOPATH := $(shell go env GOPATH)
DOCKER_BUILD=$(shell pwd)/.docker_build
DOCKER_CMD=$(DOCKER_BUILD)/go-getting-started

.PHONY: all
all: test build

$(GOPATH)/bin/dep:
	@go get github.com/golang/dep/cmd/dep

.PHONY: dep
dep: $(GOPATH)/bin/dep
	@dep ensure -v

.PHONY: dep-vendor-only
dep-vendor-only: $(GOPATH)/bin/dep
	@dep ensure -v -vendor-only

.PHONY: build
build: dep-vendor-only 
	CGO_ENABLED=0 go build -o bin/server \
        -ldflags "-X main.version=$(VERSION) -X main.serviceName=$(SERVICE_NAME)" \
		github.com/PhysicsEngine/$(SERVICE_NAME)
.PHONY: format
format:
	@gofmt -w ./..

.PHONY: check-format
check-format:
	@test -z `go fmt ./... | tee /dev/stderr`

.PHONY: test
test: dep-vendor-only check-format
	@go test -v -race ./...

.PHONY: run
run:
	@$(GO_ENV) go run \
		-ldflags "-X main.version=$(VERSION)" \
		main.go

.PHONY: coverage
coverage:
	@go test -race -coverpkg=./... -coverprofile=coverage.txt ./...

$(DOCKER_CMD): clean
	mkdir -p $(DOCKER_BUILD)
	$(GO_BUILD_ENV) go build -v -o $(DOCKER_CMD) .

.PHONY: clean
clean:
	rm -rf $(DOCKER_BUILD)

.PHONY: heroku
heroku: $(DOCKER_CMD)
	heroku container:push web
