package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"log"
)

type MyEvent struct {
	Name string `json:"name"`
}

func HandleRequest(ctx context.Context, name MyEvent) (string, error) {
	return fmt.Sprintf("Hello %s, from CodeCommit", name.Name), nil
}

func main() {
	log.Print("Handling request")
	lambda.Start(HandleRequest)
}
