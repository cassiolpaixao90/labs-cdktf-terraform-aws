package main

import (
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(hello)
}

func hello() (string, error) {
	return "Receba o mini missil aleatorio", nil
}
