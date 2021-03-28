build: deps
	mkdir -p functions
	GOBIN=${PWD}/functions go install ./...
	npm run build
.PHONY: build

build-local: deps
	go build -o tools/api-server tools/api-server.go
	npm run build
.PHONY: build-local

deps:
	@GO111MODULE=on go mod download
	@GO111MODULE=on go mod vendor
	npm install
.PHONY: deps

serve-api:
	tools/api-server
.PHONY: serve-api
