package main

import (
	"encoding/json"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"github.com/garrettreed/garrettreed.info/api/aggregate"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	siteData, siteDataErr := aggregate.GetAggregateData()
	if siteDataErr != nil {
		log.Fatal(siteDataErr)
	}

	siteDataJSON, jsonErr := json.Marshal(siteData)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(siteDataJSON),
		Headers:    map[string]string{"content-type": "application/json"},
	}, nil
}

func main() {
	lambda.Start(handler)
}
