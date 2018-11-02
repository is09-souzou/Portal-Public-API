package model

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// GetSVC get DynamoDB SVC
func GetSVC() (*dynamodb.DynamoDB, error) {
	session, err := session.NewSession(
		&aws.Config{Region: aws.String("ap-northeast-1")},
	)

	if err != nil {
		return nil, err
	}

	svc := dynamodb.New(session)

	return svc, nil
}

// User DynamoDB Resut User Struct
type User struct {
	ID          string
	Email       *string
	DisplayName string
	Career      *string
	AvatarURI   *string
	Message     *string
}

// Work DynamoDB Result Work Struct
type Work struct {
	ID          string
	UserID      string
	Title       string
	Tags        *[]string
	ImageURL    *string
	Description string
	CreatedAt   string
}
