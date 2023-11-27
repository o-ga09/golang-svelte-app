package handlers

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
) 

func Handler(ctx context.Context) (events.APIGatewayProxyResponse, error){
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body: "Hello World",
	}, nil
}