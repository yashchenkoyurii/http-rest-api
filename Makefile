.PHONY: build
build: 
	go build -v ./cmd/apiserver

.PHONY: test
test:
	go test -v -race -timeout 30s ./...
.PHONY: build-docker
build-docker:
	docker build -t apiserver:latest .

.DEFAULT_GOAL := build