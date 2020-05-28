package dynamodb

import (
	"context"
	"fmt"
	"noter/src/internal/notes"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type Adapter struct {
	tableName string
	client    *dynamodb.DynamoDB
}

//todo remove when list will be implemented
var list = []*notes.Note{
	{
		NoteId: "214e2wdq",
		Text:   "Some note",
	},
}

func (a *Adapter) PersistNote(ctx context.Context, note *notes.Note) (*notes.Note, error) {
	itemMap, err := dynamodbattribute.MarshalMap(note)
	if err != nil {
		return nil, fmt.Errorf("marshal item: %w", err)
	}

	input := &dynamodb.PutItemInput{
		Item:      itemMap,
		TableName: &a.tableName,
	}

	if _, err := a.client.PutItemWithContext(ctx, input); err != nil {
		return nil, fmt.Errorf("put item: %w", err)
	}

	return note, nil
}

func (a *Adapter) GetSingleNote(ctx context.Context, noteId string) (*notes.Note, error) {
	input := dynamodb.QueryInput{
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":nodeId": {S: aws.String(noteId)},
		},
		KeyConditionExpression: aws.String("noteId = :nodeId"),

		TableName: aws.String(a.tableName),
	}

	result, err := a.client.QueryWithContext(ctx, &input)
	if err != nil {
		return nil, fmt.Errorf("query latest note: %w", err)
	}

	if len(result.Items) == 0 {
		return nil, fmt.Errorf("note with id %v not found", noteId)
	}

	var item notes.Note
	if err := dynamodbattribute.UnmarshalMap(result.Items[0], &item); err != nil {
		return nil, fmt.Errorf("unmarshal DynamoDB item: %w", err)
	}

	return &item, nil
}

func (d *Adapter) ListNotes(ctx context.Context) ([]*notes.Note, error) {
	//todo implement list
	return list, nil
}
