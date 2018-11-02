package router

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/is09-souzou/Portal-Public-Api/src/handler"
)

// Payload request Payload struct
type Payload struct {
	Arguments json.RawMessage `json:"arguments"`
}

// Router Routing By Field
func Router(context context.Context, payload Payload) (events.APIGatewayProxyResponse, error) {
	var p handler.ListWork
	json.Unmarshal(payload.Arguments, &p)
	var data, err = handler.ListWorkHandle(p)
	binaryResult, err := json.Marshal(data)
	if err != nil {
		// 500系のResponse返す処理
		return events.APIGatewayProxyResponse{
			Body:       "Json type conversion failed",
			StatusCode: 500,
		}, nil
	}
	return events.APIGatewayProxyResponse{
		Headers: map[string]string{
			"Content-Type":                 "application/json",
			"Access-Control-Allow-Headers": "Content-Type,X-Amz-Date,Authorization,X-Api-Key,X-Amz-Security-Token",
			"Access-Control-Allow-Methods": "DELETE,GET,HEAD,OPTIONS,PATCH,POST,PUT",
			"Access-Control-Allow-Origin":  "*",
		},
		Body:       string(binaryResult),
		StatusCode: 200,
	}, nil
}
