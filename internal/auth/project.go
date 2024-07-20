package auth

import (
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type Project struct {
	Title     string `json:"title" dynamodbav:"Title"`
	AccessKey string `json:"access_key" dynamodbav:"AccessKey"`
	Secret    string `json:"secret" dynamodbav:"Secret"`
}

type Repository interface {
	FindOneByAccessKey(accessKey string) (*Project, error)
	Create(request CreateProjectRequest) error
}

func (project Project) GetKey() (map[string]types.AttributeValue, error) {
	accessKey, err := attributevalue.Marshal(project.AccessKey)
	if err != nil {
		return nil, err
	}

	return map[string]types.AttributeValue{"AccessKey": accessKey}, nil
}
