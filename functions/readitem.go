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

// var accesskey string = "AKIASA45Q7S6M3LEBBL6"
// var reagion string = "ap-northeast-1"
// var secret string = "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"
var tableName string = "Movies"

// var Title string = "kgf3"
var movieid string = "2010"

// This ReadItemId function is used to read the specfic item from an dynamo db
func ReadingItemid(w http.ResponseWriter, r *http.Request) {
	sess, _ := session.NewSession(&aws.Config{
		Region:      aws.String(reagion),
		Credentials: credentials.NewStaticCredentials(accesskey, secret, ""),
	})
	fmt.Println(sess.Config.Credentials.Get())
	svc := dynamodb.New(sess, &aws.Config{Endpoint: aws.String("")})

	// tableName := "Movies"
	// Title := "kgf3"
	// movieid := "2010"

	result, err := svc.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"Movieid": {
				N: aws.String(movieid),
			},
			"Title": {
				S: aws.String(Title),
			},
		},
	})
	fmt.Println(result)
	if err != nil {
		log.Fatalf("Got error calling GetItem: %s", err)
	}
	//id, err := strconv.Atoi(movieid)

	if Title == "" {
		log.Fatalf("No item: %s", err)
	}

	item := Item{}

	err = dynamodbattribute.UnmarshalMap(result.Item, &item)
	if err != nil {
		panic(fmt.Sprintf("Failed to unmarshal Record, %v", err))
	}

	fmt.Println("Found item:")
	fmt.Println("Id:  ", item.Movieid)
	fmt.Println("Title: ", item.Title)
	fmt.Println("Hero:  ", item.Hero)

}
