package main



import (

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"fmt"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/service/dynamodb"


)

func main()  {
	//createTab()
	//putItem()
	//getItem()
	scan()
}

func scan(){
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("cn-northwest-1"),
		Credentials: credentials.NewStaticCredentials("AKIAPFDFDDUVOMJRDFTA",
			"bGhBzgVXQQfRgm+U06fVLu44WEVZVdL285wnyXPz", ""),
	})

	svc := dynamodb.New(sess)
	input := &dynamodb.ScanInput{
		//ExpressionAttributeNames: map[string]*string{
		//	"AT": aws.String("AlbumTitle"),
		//	"ST": aws.String("SongTitle"),
		//},
		//ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
		//	":a": {
		//		S: aws.String("No One You Know"),
		//	},
		//},
		//FilterExpression:     aws.String("Artist = :a"),
		ProjectionExpression: aws.String("imageurl"),
		//ProjectionExpression: ,
		
		TableName:            aws.String("MyImage"),
	}

	result, err := svc.Scan(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case dynamodb.ErrCodeProvisionedThroughputExceededException:
				fmt.Println(dynamodb.ErrCodeProvisionedThroughputExceededException, aerr.Error())
			case dynamodb.ErrCodeResourceNotFoundException:
				fmt.Println(dynamodb.ErrCodeResourceNotFoundException, aerr.Error())
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



	for k, v := range result.Items {
		fmt.Println(k, v)

		for k1, v1 := range v{
			fmt.Println(k1, v1)

		}
	}
}


func getItem(){
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("cn-northwest-1"),
		Credentials: credentials.NewStaticCredentials("AKIAPFDFDDUVOMJRDFTA",
			"bGhBzgVXQQfRgm+U06fVLu44WEVZVdL285wnyXPz", ""),
	})

	svc := dynamodb.New(sess)
	//input := &dynamodb.GetItemInput{
	//	Key: map[string]*dynamodb.AttributeValue{
	//		"Myid": {
	//			S: aws.String("1"),
	//		},
	//	},
	//	TableName: aws.String("MyImage"),
	//}
	input := &dynamodb.QueryInput{
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":v1": {
				S: aws.String("https://mubucket.s3.cn-northwest-1.amazonaws.com.cn/dlrb.png"),
			},
		},
		KeyConditionExpression: aws.String("imageurl = :v1"),
		//ProjectionExpression:   aws.String("Myid"),
		TableName:              aws.String("MyImage"),
	}

	//result, err := svc.GetItem(input)
	result, err := svc.Query(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case dynamodb.ErrCodeProvisionedThroughputExceededException:
				fmt.Println(dynamodb.ErrCodeProvisionedThroughputExceededException, aerr.Error())
			case dynamodb.ErrCodeResourceNotFoundException:
				fmt.Println(dynamodb.ErrCodeResourceNotFoundException, aerr.Error())
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

func putItem(){

	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("cn-northwest-1"),
		Credentials: credentials.NewStaticCredentials("AKIAPFDFDDUVOMJRDFTA",
			"bGhBzgVXQQfRgm+U06fVLu44WEVZVdL285wnyXPz", ""),
	})

	svc := dynamodb.New(sess)
	input := &dynamodb.PutItemInput{
		Item: map[string]*dynamodb.AttributeValue{
			"imageurl": {
				S: aws.String("https://mubucket.s3.cn-northwest-1.amazonaws.com.cn/black.png"),
			},
			"Myid": {
				S: aws.String("2"),
			},
		},
		ReturnConsumedCapacity: aws.String("TOTAL"),
		TableName:              aws.String("MyImage"),
	}

	result, err := svc.PutItem(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case dynamodb.ErrCodeConditionalCheckFailedException:
				fmt.Println(dynamodb.ErrCodeConditionalCheckFailedException, aerr.Error())
			case dynamodb.ErrCodeProvisionedThroughputExceededException:
				fmt.Println(dynamodb.ErrCodeProvisionedThroughputExceededException, aerr.Error())
			case dynamodb.ErrCodeResourceNotFoundException:
				fmt.Println(dynamodb.ErrCodeResourceNotFoundException, aerr.Error())
			case dynamodb.ErrCodeItemCollectionSizeLimitExceededException:
				fmt.Println(dynamodb.ErrCodeItemCollectionSizeLimitExceededException, aerr.Error())
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

func createTab(){
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
				AttributeName: aws.String("Myid"),
				AttributeType: aws.String("S"),
			},
		},
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("Imageurl"),
				KeyType:       aws.String("HASH"),
			},
			{
				AttributeName: aws.String("Myid"),
				KeyType:       aws.String("RANGE"),
			},
		},
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(5),
			WriteCapacityUnits: aws.Int64(5),
		},
		TableName: aws.String("MyImage"),
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