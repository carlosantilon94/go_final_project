package main

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"log"
)

func GetData(ctx context.Context, id string) (result Item) {
	// Load the Shared AWS Configuration (~/.aws/config)
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		log.Fatal(err)
	}

	if err != nil {
		panic(err)
	}

	svc := dynamodb.NewFromConfig(cfg)
	out, err := svc.GetItem(context.TODO(), &dynamodb.GetItemInput{
		TableName: aws.String("MyTable"),
		Key: map[string]types.AttributeValue {
			"Id": &types.AttributeValueMemberS{Value: id},
		},
	})

	if err != nil {
		panic(err)
	}
	item := Item{
		Id:   "",
		Name: "",
	}
	err = attributevalue.UnmarshalMap(out.Item,&item)
	if err != nil {
		panic(err)
	}
	return item

}
