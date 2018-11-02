package model

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
)

// WorkTableName DynamoDB Work Table Name
const WorkTableName = "portal-works"

// ScanWorkListResult result ScanWorkList
type ScanWorkListResult struct {
	Items             []Work
	ExclusiveStartKey *string
}

// ExclusiveStartKey ExclusiveStartKey for pagination
type ExclusiveStartKey struct {
	ID        string `json:"id"`
	CreatedAt string `json:"createdAt"`
}

// ScanWorkList Scan work list from DynamoDB
func ScanWorkList(svc *dynamodb.DynamoDB, limit int64, exclusiveStartKey *string) (ScanWorkListResult, error) {

	params := &dynamodb.QueryInput{
		Limit:     &limit,
		TableName: aws.String(WorkTableName),
		KeyConditions: map[string]*dynamodb.Condition{
			"system": {
				ComparisonOperator: aws.String("EQ"),
				AttributeValueList: []*dynamodb.AttributeValue{
					{
						S: aws.String("work"),
					},
				},
			},
		},
		IndexName:        aws.String("system-createdAt-index"),
		ScanIndexForward: aws.Bool(false),
	}

	if exclusiveStartKey != nil {

		jsonBytes := ([]byte)(*exclusiveStartKey)

		var key ExclusiveStartKey
		json.Unmarshal(jsonBytes, &key)

		params.ExclusiveStartKey = map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(key.ID),
			},
			"createdAt": {
				S: aws.String(key.CreatedAt),
			},
			"system": {
				S: aws.String("work"),
			},
		}
	}

	result, err := svc.Query(params)

	if err != nil {
		return ScanWorkListResult{}, err
	}

	items := []Work{}

	for _, i := range result.Items {
		item := Work{}

		err := dynamodbattribute.UnmarshalMap(i, &item)

		if err != nil {
			fmt.Println("Got error unmarshalling:")
			fmt.Println(err.Error())
			return ScanWorkListResult{}, err
		}

		items = append(items, item)
	}

	var respExclusiveStartKey *string
	if result.LastEvaluatedKey != nil {

		exclusiveStartKey := ExclusiveStartKey{
			ID:        *result.LastEvaluatedKey["id"].S,
			CreatedAt: *result.LastEvaluatedKey["createdAt"].S,
		}
		byteExclusiveStartKey, err := json.Marshal(exclusiveStartKey)

		if err != nil {
			fmt.Println("Got error json Marshal exclusiveStartKey")
			fmt.Println(err.Error())
			return ScanWorkListResult{}, err
		}

		stringExclusiveStartKey := string(byteExclusiveStartKey)
		respExclusiveStartKey = &stringExclusiveStartKey
	}

	return ScanWorkListResult{items, respExclusiveStartKey}, nil
}

// ScanWorkListByTags Scan work list By Tags from DynamoDB
func ScanWorkListByTags(svc *dynamodb.DynamoDB, limit int64, exclusiveStartKey *string, tags []string) (ScanWorkListResult, error) {

	params := &dynamodb.QueryInput{
		KeyConditions: map[string]*dynamodb.Condition{
			"system": {
				ComparisonOperator: aws.String("EQ"),
				AttributeValueList: []*dynamodb.AttributeValue{
					{
						S: aws.String("work"),
					},
				},
			},
		},
		Limit:            &limit,
		TableName:        aws.String(WorkTableName),
		IndexName:        aws.String("system-createdAt-index"),
		ScanIndexForward: aws.Bool(false),
	}

	if len(tags) != 0 {
		var filt expression.ConditionBuilder

		for i, x := range tags {
			if i == 0 {
				filt = expression.Name("tags").Contains(x)
			} else {
				filt = filt.And(expression.Name("tags").Contains(x))
			}
		}

		expr, err := expression.NewBuilder().WithFilter(filt).Build()

		if err != nil {
			return ScanWorkListResult{}, err
		}

		params.ExpressionAttributeNames = expr.Names()
		params.ExpressionAttributeValues = expr.Values()
		params.FilterExpression = expr.Filter()
	}

	if exclusiveStartKey != nil {

		jsonBytes := ([]byte)(*exclusiveStartKey)

		var key ExclusiveStartKey
		json.Unmarshal(jsonBytes, &key)

		params.ExclusiveStartKey = map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(key.ID),
			},
			"createdAt": {
				S: aws.String(key.CreatedAt),
			},
			"system": {
				S: aws.String("work"),
			},
		}
	}

	result, err := svc.Query(params)

	if err != nil {
		return ScanWorkListResult{}, err
	}

	items := []Work{}

	for _, i := range result.Items {
		item := Work{}

		err := dynamodbattribute.UnmarshalMap(i, &item)

		if err != nil {
			fmt.Println("Got error unmarshalling:")
			fmt.Println(err.Error())
			return ScanWorkListResult{}, err
		}

		items = append(items, item)
	}

	var respExclusiveStartKey *string
	if result.LastEvaluatedKey != nil {

		exclusiveStartKey := ExclusiveStartKey{
			ID:        *result.LastEvaluatedKey["id"].S,
			CreatedAt: *result.LastEvaluatedKey["createdAt"].S,
		}
		byteExclusiveStartKey, err := json.Marshal(exclusiveStartKey)

		if err != nil {
			fmt.Println("Got error json Marshal exclusiveStartKey")
			fmt.Println(err.Error())
			return ScanWorkListResult{}, err
		}

		stringExclusiveStartKey := string(byteExclusiveStartKey)
		respExclusiveStartKey = &stringExclusiveStartKey
	}

	return ScanWorkListResult{items, respExclusiveStartKey}, nil
}

// ScanWorkListByUserID Scan work list By User ID from DynamoDB
func ScanWorkListByUserID(svc *dynamodb.DynamoDB, limit int64, exclusiveStartKey *string, userID string) (ScanWorkListResult, error) {

	filt := expression.Name("userId").Contains(userID)
	expr, err := expression.NewBuilder().WithFilter(filt).Build()

	if err != nil {
		return ScanWorkListResult{}, err
	}

	params := &dynamodb.QueryInput{
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		FilterExpression:          expr.Filter(),
		KeyConditions: map[string]*dynamodb.Condition{
			"system": {
				ComparisonOperator: aws.String("EQ"),
				AttributeValueList: []*dynamodb.AttributeValue{
					{
						S: aws.String("work"),
					},
				},
			},
		},
		Limit:            &limit,
		TableName:        aws.String(WorkTableName),
		IndexName:        aws.String("system-createdAt-index"),
		ScanIndexForward: aws.Bool(false),
	}

	if exclusiveStartKey != nil {

		jsonBytes := ([]byte)(*exclusiveStartKey)

		var key ExclusiveStartKey
		json.Unmarshal(jsonBytes, &key)

		params.ExclusiveStartKey = map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(key.ID),
			},
			"createdAt": {
				S: aws.String(key.CreatedAt),
			},
			"system": {
				S: aws.String("work"),
			},
		}
	}

	result, err := svc.Query(params)

	if err != nil {
		return ScanWorkListResult{}, err
	}

	items := []Work{}

	for _, i := range result.Items {
		item := Work{}

		err := dynamodbattribute.UnmarshalMap(i, &item)

		if err != nil {
			fmt.Println("Got error unmarshalling:")
			fmt.Println(err.Error())
			return ScanWorkListResult{}, err
		}

		items = append(items, item)
	}

	var respExclusiveStartKey *string
	if result.LastEvaluatedKey != nil {

		exclusiveStartKey := ExclusiveStartKey{
			ID:        *result.LastEvaluatedKey["id"].S,
			CreatedAt: *result.LastEvaluatedKey["createdAt"].S,
		}
		byteExclusiveStartKey, err := json.Marshal(exclusiveStartKey)

		if err != nil {
			fmt.Println("Got error json Marshal exclusiveStartKey")
			fmt.Println(err.Error())
			return ScanWorkListResult{}, err
		}

		stringExclusiveStartKey := string(byteExclusiveStartKey)
		respExclusiveStartKey = &stringExclusiveStartKey
	}

	return ScanWorkListResult{items, respExclusiveStartKey}, nil
}
