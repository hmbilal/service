package db

import (
	"context"
	"github.com/hmbilal/gofiber-start/pkg/config"

	awsConfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func NewDynamoDBConnection(cfg *config.Configuration) *dynamodb.Client {
	awsCfg, err := awsConfig.LoadDefaultConfig(context.Background())
	if err != nil {
		panic(err)
	}

	// DynamoDBLocal can be resolved only by endpoint.
	// https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/DynamoDBLocal.html
	// quote: "When you're ready to deploy your application in production, you remove the local endpoint in the code,
	// and then it points to the DynamoDB web service."
	if cfg.Environment == "local" || cfg.Environment == "test" {
		return dynamodb.NewFromConfig(awsCfg, func(o *dynamodb.Options) {
			o.BaseEndpoint = &cfg.DB.DynamoDB.Endpoint
		})
	}

	return dynamodb.NewFromConfig(awsCfg)
}
