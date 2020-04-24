package play

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

func MapToAttrVal() {

	m := map[int]string{
		1: "Acky",
		2: "Becky",
		3: "Cindy",
		4: "",
	}
	av, err := dynamodbattribute.MarshalMap(m)
	if err != nil {
		panic(err)
	}
	fmt.Println("attribute value: ", av)
	for k, v := range av {
		r := "nil"
		if v.NULL != nil {
			if *(v.NULL) {
				r = "true"
			} else {
				r = "false"
			}
		}
		fmt.Printf("key: %v, value: %v\n", k, r)
	}
}

func StructToAttrVal() {
	p := person{
		"Joe",
		50,
	}
	av, err := dynamodbattribute.MarshalMap(p)
	if err != nil {
		panic(err)
	}
	fmt.Println("struct person: ", av)
}

type person struct {
	Name string
	Age  int
}

func AttrValToMap() {
	av := map[string]*dynamodb.AttributeValue{
		"CatalogCode": {
			S: aws.String("c000"),
		},
		"EnabledOn": {
			S: aws.String("20200310"),
		},
	}
	m := make(map[string]*interface{})
	err := dynamodbattribute.UnmarshalMap(av, &m)
	if err != nil {
		panic(err)
	}
	for k, v := range m {
		fmt.Printf("key: '%v', value: '%v'\n", k, *v)
	}

}
