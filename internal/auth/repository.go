package auth

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/hmbilal/gofiber-start/pkg/config"
	"github.com/hmbilal/gofiber-start/pkg/db"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type projectRepository struct {
	db.DynamoDBRepository
	Cfg *config.Configuration
}

func NewDynamoDBRepository(
	dynamodbClient *dynamodb.Client,
	cfg *config.Configuration,
) Repository {
	return &projectRepository{
		DynamoDBRepository: *db.NewDynamoDBRepository(dynamodbClient),
		Cfg:                cfg,
	}
}

func (r *projectRepository) FindOneByAccessKey(accessKey string) (*Project, error) {
	var project Project
	project.AccessKey = accessKey

	key, err := project.GetKey()
	if err != nil {
		return nil, err
	}

	input := &dynamodb.GetItemInput{
		TableName:                aws.String(r.Cfg.DB.DynamoDB.TableNames.Projects),
		Key:                      key,
		AttributesToGet:          nil,
		ConsistentRead:           nil,
		ExpressionAttributeNames: nil,
		ProjectionExpression:     nil,
		ReturnConsumedCapacity:   "",
	}

	item, err := r.GetItem(input)
	if err != nil {
		return nil, err
	}

	err = attributevalue.UnmarshalMap(item, &project)
	if err != nil {
		return nil, err
	}

	return &project, nil
}

func (r *projectRepository) Create(request CreateProjectRequest) error {
	item, err := attributevalue.MarshalMap(Project(request))
	if err != nil {
		return err
	}

	input := &dynamodb.PutItemInput{
		Item:                        item,
		TableName:                   aws.String(r.Cfg.DB.DynamoDB.TableNames.Projects),
		ConditionExpression:         nil,
		ConditionalOperator:         "",
		Expected:                    nil,
		ExpressionAttributeNames:    nil,
		ExpressionAttributeValues:   nil,
		ReturnConsumedCapacity:      "",
		ReturnItemCollectionMetrics: "",
		ReturnValues:                "",
	}

	return r.SaveItem(input)
}
