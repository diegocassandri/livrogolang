package dbexample

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

// func GetItems(client *dynamodb.DynamoDB) (*dynamodb.GetItemOutput, error) {
// 	return client.GetItem(&dynamodb.GetItemInput{
// 		AttributesToGet: []*string{aws.String("bar")},
// 		Key: map[string]*dynamodb.AttributeValue{
// 			"foo": {
// 				BOOL: aws.Bool(true),
// 			},
// 		},
// 	})
// }

// Usando a interface dynamodbiface.DynamoDBAPI, o problema é que ela possui muitos métodos
// Como satisfazer a interface sem precisar definir todos os métodos?
func GetItems(client dynamodbiface.DynamoDBAPI) (*dynamodb.GetItemOutput, error) {
	return client.GetItem(&dynamodb.GetItemInput{
		AttributesToGet: []*string{aws.String("bar")},
		Key: map[string]*dynamodb.AttributeValue{
			"foo": {
				BOOL: aws.Bool(true),
			},
		},
	})
}
