package container

import (
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/hmbilal/gofiber-start/pkg/config"
	"github.com/hmbilal/gofiber-start/pkg/db"
)

type Container struct {
	Cfg *config.Configuration
	Db  *dynamodb.Client
}

func NewContainer(
	configFile *string,
) *Container {
	cfg := config.NewConfiguration(configFile)
	dynamo := db.NewDynamoDBConnection(cfg)

	return &Container{
		Cfg: cfg,
		Db:  dynamo,
	}
}
