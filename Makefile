#!/bin/bash

build: binary

binary:
	@echo "building binary..."
	@go build -o vektor ./main.go


clean:
	@echo "cleaning..."
	@rm -rf vendor
	@rm -f go.sum


install:
	@echo "Installing dependencies..."
	@rm -rf vendor
	@rm -f go.sum
	@go mod tidy && go mod download && go mod vendor

start:
	@go run main.go
