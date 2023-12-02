package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	handlers "github.com/o-ga09/golang-svelte-app/handlers/counter"
)

func main() {
	lambda.Start(handlers.HealthCheck)
}
