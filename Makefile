build:
	rm -rf functions
	mkdir functions
	# Must manually list dependencies instead of `go get ./api/...`
	# because netflify symlinks the repo to gopath without .git, which breaks
	# `go get` when using project-level imports, even with the `-f` flag set.
	go get github.com/aws/aws-lambda-go/events
	go get github.com/aws/aws-lambda-go/lambda
	go build -o functions/main api/main.go
	npm run build

build-local:
	go get ./api/...
	go build -o tools/api-server tools/api-server.go
	npm run build

serve-api:
	tools/api-server

serve-static:
	npm run start
