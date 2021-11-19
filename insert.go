package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
)

func InsertData(ctx context.Context, item Item) {
	cfg, err := config.LoadDefaultConfig(ctx, func(o *config.LoadOptions) error {
		o.Region = "us-east-1"
		return nil
	})
	if err != nil {
		panic(err)
	}
	data, err := attributevalue.MarshalMap(item)
	svc := dynamodb.NewFromConfig(cfg)
	out, err := svc.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: aws.String("MyTable"),
		Item: data,
	})

	if err != nil {
		panic(err)
	}

	fmt.Println(out.Attributes)
}