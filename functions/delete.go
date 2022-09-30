package functions

import (
	"fmt"
	"log"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// this is a hard delete function
func DeleteItem(w http.ResponseWriter, r *http.Request) {

	sess, _ := session.NewSession(&aws.Config{
		Region:      aws.String(reagion),
		Credentials: credentials.NewStaticCredentials(accesskey, secret, ""),
	})
	fmt.Println(sess.Config.Credentials.Get())
	svc := dynamodb.New(sess, &aws.Config{Endpoint: aws.String("")})

	// tableName := "Movies"
	movieName := "xyz"
	// movieid := "2001"

	input := &dynamodb.DeleteItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"Movieid": {
				N: aws.String(movieid),
			},
			"Title": {
				S: aws.String(movieName),
			},
		},
		TableName: aws.String(tableName),
	}
	_, err := svc.DeleteItem(input)
	if err != nil {
		log.Fatalf("Got error calling DeleteItem:%s", err)
	}
	fmt.Println("Deleted'" + movieName + "'(" + movieid + ")from table" + tableName)
}
