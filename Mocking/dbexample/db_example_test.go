package dbexample_test

import (
	"errors"
	"livrogolang/Mocking/dbexample"
	"reflect"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

// Embed the interface in a custum type
type mockDynamoDBClient struct {
	dynamodbiface.DynamoDBAPI
	output *dynamodb.GetItemOutput
	err    error
}

func (m *mockDynamoDBClient) GetItem(_ *dynamodb.GetItemInput) (*dynamodb.GetItemOutput, error) {
	return m.output, m.err
}

func TestGetItems(t *testing.T) {
	exampleErr := errors.New("uh oh")
	exampleResponse := &dynamodb.GetItemOutput{
		Item: map[string]*dynamodb.AttributeValue{
			"foo": {
				BOOL: aws.Bool(false),
			},
		},
	}

	subtest := []struct {
		name     string
		client   *mockDynamoDBClient
		response *dynamodb.GetItemOutput
		err      error
	}{
		{
			name: "happy-path",
			client: &mockDynamoDBClient{
				output: exampleResponse,
				err:    nil,
			},
			response: exampleResponse,
			err:      nil,
		},
		{
			name: "error-from-call",
			client: &mockDynamoDBClient{
				err: exampleErr,
			},
			err: exampleErr,
		},
	}
	for _, tc := range subtest {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := dbexample.GetItems(tc.client)
			if !reflect.DeepEqual(actual, tc.response) {
				t.Errorf("expected response %v, got %v", tc.response, actual)
			}
			if !errors.Is(err, tc.err) {
				t.Errorf("expected error %v, got %v", tc.err, err)
			}
		})
	}
}
