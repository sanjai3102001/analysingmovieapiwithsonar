package functions

import (
	"fmt"
	"log"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

var accesskey string = "AKIASA45Q7S6M3LEBBL6"
var reagion string = "ap-northeast-1"
var secret string = "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"

// var tableName string = "Movies"
var Title string = "kgf3"

// var movieid string = "2010"

// This CreateItem function id used to create a item in a dynamo db
func CreateItem(w http.ResponseWriter, r *http.Request) {

	sess, _ := session.NewSession(&aws.Config{
		Region:      aws.String(reagion),
		Credentials: credentials.NewStaticCredentials(accesskey, secret, ""),
	})
	fmt.Println(sess.Config.Credentials.Get())
	svc := dynamodb.New(sess, &aws.Config{Endpoint: aws.String("")})

	item := map[string]interface{}{
		"Movieid":  2022,
		"Title":    "kgf2",
		"Hero":     "yash",
		"isactive": true,
	}

	av, err := dynamodbattribute.MarshalMap(item)
	if err != nil {
		log.Fatalf("Got error marshalling new movie item: %s", err)
	}

	tableName := "Movies"

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(tableName),
	}

	_, err = svc.PutItem(input)
	if err != nil {
		log.Fatalf("Got error calling PutItem: %s", err)
	}

	// year := strconv.Itoa(item["Movieid"].(int))

	// fmt.Println("Successfully added '" + item["Title"].(string) + "' (" + year + ") to table " + tableName)
}

type Item struct {
	Movieid int
	Title   string
	Hero    string
}
