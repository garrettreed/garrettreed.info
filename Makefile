build:
	rm -rf functions
	mkdir functions
	go get github.com/aws/aws-lambda-go/events
	go get github.com/aws/aws-lambda-go/lambda
	go build -o functions/main api/main.go
