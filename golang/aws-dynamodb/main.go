package main

import (
	"context"
	"log/slog"
	"os"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func main() {
	log := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{})).With("module", "aws-dynamodb")

	ctx := context.Background()
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}

	client := dynamodb.NewFromConfig(cfg)

	listTablesOutput, err := client.ListTables(ctx, &dynamodb.ListTablesInput{})
	if err != nil {
		log.Error("ListTables failed", "err", err.Error())
		os.Exit(1)
		return
	}
	log.Info("ListTablesOutput", "TableNames", listTablesOutput.TableNames)

	putItemOutput, err := client.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: aws.String("todo"),
		Item: map[string]types.AttributeValue{
			"id":   &types.AttributeValueMemberS{Value: "1"},
			"date": &types.AttributeValueMemberS{Value: time.Now().Format(time.RFC3339)},
		},
		ConditionExpression: aws.String("attribute_not_exists(id)"),
	})
	if err != nil {
		log.Error("PutItem failed", "err", err.Error())
		os.Exit(1)
		return
	}
	log.Info("PutItemOutput", "ResultMetadata", putItemOutput.ResultMetadata)

	getItemOutput, err := client.GetItem(ctx, &dynamodb.GetItemInput{
		TableName: aws.String("todo"),
		Key: map[string]types.AttributeValue{
			"id": &types.AttributeValueMemberS{Value: "1"},
		},
	})
	if err != nil {
		log.Error("GetItem failed", "err", err.Error())
		os.Exit(1)
		return
	}
	log.Info("GetItemOutput", "item", getItemOutput.Item)

	deleteItemOutput, err := client.DeleteItem(ctx, &dynamodb.DeleteItemInput{
		TableName: aws.String("todo"),
		Key: map[string]types.AttributeValue{
			"id": &types.AttributeValueMemberS{Value: "1"},
		},
	})
	if err != nil {
		log.Error("DeleteItem failed", "err", err.Error())
		os.Exit(1)
		return
	}
	log.Info("DeleteItemOutput", "ResultMetadata", deleteItemOutput.ResultMetadata)
}
