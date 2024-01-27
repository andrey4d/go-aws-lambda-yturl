package main

import (
	"github.com/andrey4d/go-aws-lambda-yturl/internal/handlers"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {

	lambda.Start(handlers.HandleRequest)
}
