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

// This Softdelete function id used to delete a item in a dynamo db
func Softdelete(w http.ResponseWriter, r *http.Request) {

	sess, _ := session.NewSession(&aws.Config{
		Region:      aws.String(reagion),
		Credentials: credentials.NewStaticCredentials(accesskey, secret, ""),
	})
	fmt.Println(sess.Config.Credentials.Get())
	svc := dynamodb.New(sess, &aws.Config{Endpoint: aws.String("")})

	// tableName := "Movies"
	// Title := "kgf2"
	Movieid := "2022"
	Isactive := false

	input := &dynamodb.UpdateItemInput{
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":r": {
				BOOL: aws.Bool(Isactive),
			},
		},
		TableName: aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"Movieid": {
				N: aws.String(Movieid),
			},
			"Title": {
				S: aws.String(Title),
			},
		},
		ReturnValues:     aws.String("UPDATED_NEW"),
		UpdateExpression: aws.String("set Isactive = :r"),
	}

	_, err := svc.UpdateItem(input)
	if err != nil {
		log.Fatalf("Got error calling UpdateItem: %s", err)
	}

}
