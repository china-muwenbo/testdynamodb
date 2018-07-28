package myUpload


import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"fmt"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"encoding/json"
	"os"
	"net/http"
)

func main() {
	//createDB()
	//putUrlItem()
	//QueryAll()
}


type urls struct {
	Imageurl string `json:imageurl`
	ImageName string `json:imagename`
}

func QueryAll(w http.ResponseWriter, r *http.Request) {
	b:=queryAll()
	http.DetectContentType(b)
	w.Write(b)
}
func queryAll() []byte {

	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("cn-northwest-1"),
		Credentials: credentials.NewStaticCredentials("AKIAPFDFDDUVOMJRDFTA",
			"bGhBzgVXQQfRgm+U06fVLu44WEVZVdL285wnyXPz", ""),
	})

	svc := dynamodb.New(sess)

	input := &dynamodb.ScanInput{
		//ProjectionExpression: aws.String("Imageurl"),
		TableName:            aws.String("MuImage"),
	}


	result, err := svc.Scan(input)
	urlss := []urls{}
	dynamodbattribute.UnmarshalListOfMaps(result.Items,&urlss)
	if err != nil {

	}
	b,_:=json.Marshal(urlss)
	os.Stdout.Write(b)
	//fmt.Println(string())

	return b
}



func putUrlItem(url string){
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("cn-northwest-1"),
		Credentials: credentials.NewStaticCredentials("AKIAPFDFDDUVOMJRDFTA",
			"bGhBzgVXQQfRgm+U06fVLu44WEVZVdL285wnyXPz", ""),
	})

	svc := dynamodb.New(sess)
	input := &dynamodb.PutItemInput{
		Item: map[string]*dynamodb.AttributeValue{
			"Imageurl": {
				S: aws.String(url),
			},
			"ImageName": {
				S: aws.String("black.png"),
			},
		},
		ReturnConsumedCapacity: aws.String("TOTAL"),
		TableName:              aws.String("MuImage"),
	}

	result, err := svc.PutItem(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case dynamodb.ErrCodeResourceInUseException:
				fmt.Println(dynamodb.ErrCodeResourceInUseException, aerr.Error())
			case dynamodb.ErrCodeLimitExceededException:
				fmt.Println(dynamodb.ErrCodeLimitExceededException, aerr.Error())
			case dynamodb.ErrCodeInternalServerError:
				fmt.Println(dynamodb.ErrCodeInternalServerError, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

func createDB(){
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("cn-northwest-1"),
		Credentials: credentials.NewStaticCredentials("AKIAPFDFDDUVOMJRDFTA",
			"bGhBzgVXQQfRgm+U06fVLu44WEVZVdL285wnyXPz", ""),
	})

	svc := dynamodb.New(sess)
	input := &dynamodb.CreateTableInput{
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("Imageurl"),
				AttributeType: aws.String("S"),
			},
			{
				AttributeName: aws.String("ImageName"),
				AttributeType: aws.String("S"),
			},
		},
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("Imageurl"),
				KeyType:       aws.String("HASH"),
			},
			{
				AttributeName: aws.String("ImageName"),
				KeyType:       aws.String("RANGE"),
			},
		},
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(5),
			WriteCapacityUnits: aws.Int64(5),
		},
		TableName: aws.String("MuImage"),
	}

	result, err := svc.CreateTable(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case dynamodb.ErrCodeResourceInUseException:
				fmt.Println(dynamodb.ErrCodeResourceInUseException, aerr.Error())
			case dynamodb.ErrCodeLimitExceededException:
				fmt.Println(dynamodb.ErrCodeLimitExceededException, aerr.Error())
			case dynamodb.ErrCodeInternalServerError:
				fmt.Println(dynamodb.ErrCodeInternalServerError, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

