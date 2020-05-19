package play

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

var sess = session.Must(session.NewSession())
var client = dynamodb.New(sess, aws.NewConfig().WithRegion("ap-northeast-1"))

func DDBGetItem() {
	input := &dynamodb.GetItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"CatalogCode": {
				S: aws.String("aaa"),
			},
			"PrimarySK": {
				S: aws.String("202004101000_S02"),
			},
		},
		ProjectionExpression: aws.String("ProcessedAt, SupplierCode"),
		TableName:            aws.String("Products"),
	}
	r, err := client.GetItem(input)
	if err != nil {
		panic(err)
	}
	fmt.Println("is nil : ", len(r.Item) == 0)
}

func DDBUpdateItem() {
	var variationLinks interface{} = ""
	var isFurniture interface{} = false
	var icon interface{} = "icon3"
	payload := map[string]*interface{}{
		"VariationLinks": &variationLinks,
		"IsFurniture":    &isFurniture,
		"Icon1":          &icon,
	}
	p := Product{
		ProcessedAt: "20200304",
		Payload:     payload,
	}
	creator := DefaultUpdateStatementCreator{}
	statement, err := creator.Create(&p)
	if err != nil {
		panic(err)
	}
	update := dynamodb.UpdateItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"CatalogCode": {
				S: aws.String("c100"),
			},
			"EnabledOn": {
				S: aws.String("2020101011"),
			},
		},
		UpdateExpression:          &statement.Expression,
		ExpressionAttributeValues: statement.AttributeValues,
		// ConditionExpression:       aws.String("attribute_exists(CatalogCode) AND attribute_exists(EnabledOn)"),
		TableName: aws.String("SimpleProduct"),
	}
	output, err := client.UpdateItem(&update)
	if err != nil {
		panic(err)
	}
	fmt.Println("update output: ", output)
}

type Product struct {
	CatalogCode  string `dynamodbav:"CatalogCode"`
	EnabledOn    string `dynamodbav:"EnabledOn"`
	SupplierCode string `dynamodbav:"SupplierCode,omitempty"`
	ProcessedAt  string `dynamodbav:"ProcessedAt"`
	Payload      map[string]*interface{}
}

type NewProduct struct {
	heap map[string]*interface{}
}

type UpdateStatementCreator interface {
	Create(prod *Product) (UpdateStatement, error)
}

type DefaultUpdateStatementCreator struct {
}

func (c DefaultUpdateStatementCreator) Create(prod *Product) (UpdateStatement, error) {
	payload := prod.Payload
	attrsToRemove := make([]string, 0, len(payload))
	for k, v := range payload {
		if v == nil {
			attrsToRemove = append(attrsToRemove, k)
		}
	}
	mpayload, err := dynamodbattribute.MarshalMap(prod.Payload)
	if err != nil {
		return UpdateStatement{}, err
	}
	attrsToUpdate := make([]string, 0, len(mpayload))
	attrsToIgnore := make([]string, 0, len(mpayload))
	for k, v := range mpayload {
		if v.NULL != nil && *(v.NULL) {
			attrsToIgnore = append(attrsToIgnore, k)
		} else {
			attrsToUpdate = append(attrsToUpdate, k)
		}
	}
	for _, k := range attrsToIgnore {
		delete(mpayload, k)
	}
	avs := map[string]*dynamodb.AttributeValue{
		":ProcessedAt": {
			S: &prod.ProcessedAt,
		},
	}
	set := "SET ProcessedAt = :ProcessedAt"
	for k, v := range mpayload {
		arg := ":" + k
		set += ", " + k + " = " + arg
		avs[arg] = v
	}
	remove := ""
	for i, k := range attrsToRemove {
		if i == 0 {
			remove += "REMOVE " + k
		} else {
			remove += ", " + k
		}
	}
	return UpdateStatement{
		Expression:      set + " " + remove,
		AttributeValues: avs,
	}, nil
}

type UpdateStatement struct {
	Expression      string
	AttributeValues map[string]*dynamodb.AttributeValue
}
