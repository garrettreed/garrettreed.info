.PHONY: build
build:
	mkdir functions
	go get ./...
	go build -o functions/main api/main.go
	npm run build

.PHONY: build-local
build-local:
	go get ./...
	go build -o tools/api-server tools/api-server.go
	npm run build

.PHONY: serve-api
serve-api:
	tools/api-server

.PHONY: serve-static
serve-static:
	npm run start
