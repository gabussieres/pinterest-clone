package resources

import (
	"fmt"

	"pinterest-clone/config"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

/*
This is a go-DynamoDB library that would ideally be separated into its own module in another repository.
*/

var svc *dynamodb.DynamoDB

func init() {
	// Initialize a session that the SDK will use to load
	// credentials from the shared credentials file ~/.aws/credentials
	// and region from the shared configuration file ~/.aws/config.
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Create DynamoDB client
	svc = dynamodb.New(sess)
}

// DynamoInterface holds methods that interact with the DynamoDB client
type DynamoInterface interface {
	GetItem(id string, tableConfig config.TableConfig) (*dynamodb.GetItemOutput, error)
	BatchGetItem(ids interface{}, tableConfig config.TableConfig) (*dynamodb.BatchGetItemOutput, error)
	Query(id string, tableConfig config.TableConfig) (*dynamodb.QueryOutput, error)
}

// Dynamo implements the DynamoInterface
type Dynamo struct{}

// GetItem gets an item from the table set in tableConfig, given an id
func (d Dynamo) GetItem(id string, tableConfig config.TableConfig) (*dynamodb.GetItemOutput, error) {
	return d.getItem(svc, id, tableConfig)
}

func (d Dynamo) getItem(dbiface dynamodbiface.DynamoDBAPI, id string, tableConfig config.TableConfig) (out *dynamodb.GetItemOutput, err error) {
	out, err = dbiface.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(tableConfig.Name),
		Key: map[string]*dynamodb.AttributeValue{
			tableConfig.PartitionKey: {
				S: aws.String(id),
			},
		},
	})
	if err != nil {
		fmt.Println(err.Error())
	}
	return
}

func createBatchGetItemKeyMap(ids interface{}, pk string) (input []map[string]*dynamodb.AttributeValue) {
	for _, id := range ids.(Pins) {
		key := map[string]*dynamodb.AttributeValue{
			pk: &dynamodb.AttributeValue{
				S: aws.String(string(id)),
			},
		}
		input = append(input, key)
	}
	return
}

// BatchGetItem gets multiple items from the table set in tableConfig, given an array if ids
func (d Dynamo) BatchGetItem(ids interface{}, tableConfig config.TableConfig) (*dynamodb.BatchGetItemOutput, error) {
	return d.batchGetItem(svc, ids, tableConfig)
}

func (d Dynamo) batchGetItem(dbiface dynamodbiface.DynamoDBAPI, ids interface{}, tableConfig config.TableConfig) (out *dynamodb.BatchGetItemOutput, err error) {
	out, err = dbiface.BatchGetItem(&dynamodb.BatchGetItemInput{
		RequestItems: map[string]*dynamodb.KeysAndAttributes{
			tableConfig.Name: {
				Keys: createBatchGetItemKeyMap(ids, tableConfig.PartitionKey),
			},
		},
	})
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case dynamodb.ErrCodeProvisionedThroughputExceededException:
				fmt.Println(dynamodb.ErrCodeProvisionedThroughputExceededException, aerr.Error())
			case dynamodb.ErrCodeResourceNotFoundException:
				fmt.Println(dynamodb.ErrCodeResourceNotFoundException, aerr.Error())
			case dynamodb.ErrCodeRequestLimitExceeded:
				fmt.Println(dynamodb.ErrCodeRequestLimitExceeded, aerr.Error())
			case dynamodb.ErrCodeInternalServerError:
				fmt.Println(dynamodb.ErrCodeInternalServerError, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			fmt.Println(err.Error())
		}
		return
	}
	return
}

// Query returns an array of items for the given id
func (d Dynamo) Query(id string, tableConfig config.TableConfig) (*dynamodb.QueryOutput, error) {
	return d.query(svc, id, tableConfig)
}

func (d Dynamo) query(dbiface dynamodbiface.DynamoDBAPI, id string, tableConfig config.TableConfig) (out *dynamodb.QueryOutput, err error) {
	input := &dynamodb.QueryInput{
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":v1": {
				S: aws.String(id),
			},
		},
		KeyConditionExpression: aws.String(fmt.Sprintf("%s = :v1", tableConfig.PartitionKey)),
		ProjectionExpression:   aws.String(tableConfig.SortKey),
		TableName:              aws.String(tableConfig.Name),
	}

	out, err = dbiface.Query(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case dynamodb.ErrCodeProvisionedThroughputExceededException:
				fmt.Println(dynamodb.ErrCodeProvisionedThroughputExceededException, aerr.Error())
			case dynamodb.ErrCodeResourceNotFoundException:
				fmt.Println(dynamodb.ErrCodeResourceNotFoundException, aerr.Error())
			case dynamodb.ErrCodeRequestLimitExceeded:
				fmt.Println(dynamodb.ErrCodeRequestLimitExceeded, aerr.Error())
			case dynamodb.ErrCodeInternalServerError:
				fmt.Println(dynamodb.ErrCodeInternalServerError, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			fmt.Println(err.Error())
		}
		return
	}
	return
}
