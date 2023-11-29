package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	di "github.com/o-ga09/golang-svelte-app/DI"
)

func main() {
	handler := di.New()
	lambda.Start(handler.GetCouter)
}
