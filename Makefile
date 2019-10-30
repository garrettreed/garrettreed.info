build:
	rm -rf functions
	mkdir functions
	go get ./...
	go build -o functions/main api/main.go
	npm run build

build-local:
	go get ./...
	go build -o tools/api-server tools/api-server.go
	npm run build

serve-api:
	tools/api-server

serve-static:
	npm run start
