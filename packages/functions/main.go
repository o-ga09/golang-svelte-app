package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/o-ga09/golang-svelte-app/handlers"
)



func main() {
	lambda.Start(handlers.Handler)
}