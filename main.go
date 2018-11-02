package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/is09-souzou/Portal-Public-Api/src/router"
)

func main() {
	lambda.Start(router.Router)
}
