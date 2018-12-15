GIT_REF := $(shell git describe --always --tag)
VERSION ?= commit-$(GIT_REF)

SERVICE_NAME := $(shell basename ${PWD})

REVIEWDOG_ARG ?= -diff="git diff master"

LINT_TOOLS=\
	golang.org/x/lint/golint \
	github.com/client9/misspell \
	github.com/kisielk/errcheck \
	honnef.co/go/tools/cmd/staticcheck

GOPATH := $(shell go env GOPATH)

.PHONY: all
all: test reviewdog

.PHONY: bootstrap-lint-tools
bootstrap-lint-tools:
	@for tool in $(LINT_TOOLS) ; do \
		echo "Installing/Updating $$tool" ; \
		go get -u $$tool; \
	done

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
	@gofmt -d ./..

.PHONY: test
test: check-format
	@go test -v -race ./...

.PHONY: run
run:
	@$(GO_ENV) go run \
		-ldflags "-X main.version=$(VERSION)" \
		main.go

.PHONY: reviewdog
reviewdog:
	@go get github.com/haya14busa/reviewdog/cmd/reviewdog
	reviewdog -conf=.reviewdog.yml $(REVIEWDOG_ARG)

.PHONY: coverage
coverage:
	@go test -race -coverpkg=./... -coverprofile=coverage.txt ./...

