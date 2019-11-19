SOURCEDIR = .
REPO=github.com/krinklesaurus
DOCKER_REGISTRY=registry.${REPO}
NAME ?= go-p2pb2b
GO_VERSION=1.13.4
GO_RUN=docker run --rm -v ${PWD}:/usr/src/myapp -w /usr/src/myapp golang:${GO_VERSION}
VERSION=$(shell git rev-parse --short HEAD)

default: clean test build

.PHONY: init
init:
	export GO111MODULE=on &&\
		${GO_RUN} go mod init ${REPO}/${NAME}


.PHONY: clean
clean:
	@if [ -f ${NAME} ] ; then rm ${NAME}; fi &&\
		rm -rf vendor/


# validate the project is correct and all necessary information is available
.PHONY: validate
validate:
	export GO111MODULE=on &&\
		${GO_RUN} /bin/bash -c "cd /usr/src/myapp && go mod tidy && go mod vendor && go mod verify"


.PHONY: test
test:
	${GO_RUN} go test -v -race -coverprofile=coverage.out $$(go list ./... | grep -v /vendor/)


.PHONY: cover
cover:
	go tool cover -html=coverage.out

.PHONY: build
build:
	${GO_RUN} go build -v
