package db

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type Model interface {
	GetKey() (map[string]types.AttributeValue, error)
}

type IDynamoDBRepository interface {
	GetItem(input *dynamodb.GetItemInput) (map[string]types.AttributeValue, error)
	SaveItem(input *dynamodb.PutItemInput) error
}

type DynamoDBRepository struct {
	client *dynamodb.Client
}

func NewDynamoDBRepository(client *dynamodb.Client) *DynamoDBRepository {
	return &DynamoDBRepository{
		client: client,
	}
}

func (r *DynamoDBRepository) GetItem(input *dynamodb.GetItemInput) (map[string]types.AttributeValue, error) {
	response, err := r.client.GetItem(context.Background(), input)
	if err != nil {
		return nil, err
	}

	return response.Item, nil
}

func (r *DynamoDBRepository) SaveItem(input *dynamodb.PutItemInput) error {
	_, err := r.client.PutItem(context.Background(), input)
	if err != nil {
		return err
	}

	return nil
}
