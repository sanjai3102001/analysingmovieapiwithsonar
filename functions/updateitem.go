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

// This is a update function for a movie api
func UpdateItems(w http.ResponseWriter, r *http.Request) {

	sess, _ := session.NewSession(&aws.Config{
		Region:      aws.String(reagion),
		Credentials: credentials.NewStaticCredentials(accesskey, secret, ""),
	})
	fmt.Println(sess.Config.Credentials.Get())
	svc := dynamodb.New(sess, &aws.Config{Endpoint: aws.String("")})

	// tableName := "Movies"
	// Title := "kgf2"
	Movieid := "2010"
	Hero := "kgdyash"

	input := &dynamodb.UpdateItemInput{
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":r": {
				S: aws.String(Hero),
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
		UpdateExpression: aws.String("set Hero = :r"),
	}

	_, err := svc.UpdateItem(input)
	if err != nil {
		log.Fatalf("Got error calling UpdateItem: %s", err)
	}

	fmt.Println("Successfully updated '" + Title + "' (" + Movieid + ") rating to " + Title)

}
