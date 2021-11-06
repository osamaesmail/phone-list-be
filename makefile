SHELL := /bin/bash

include .env.test
export $(shell sed 's/=.*//' .env.test)

.PHONY: lint
lint:
	golangci-lint run ./...

.PHONY: lint-fix
lint-fix:
	golangci-lint run ./... --fix

.PHONY: test
test:
	go test -v ./... -covermode=atomic

.PHONY: swag
swag: swag.install swag.generate

.PHONY: swag.generate
swag.generate:
	swag init -g router/router.go

.PHONY: swag.install
swag.install:
	which swag || GO111MODULE=off go get -u github.com/swaggo/swag/cmd/swag
