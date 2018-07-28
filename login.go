package main

import (
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"fmt"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"net/http"
	"io"
	"html/template"
	"log"
	"gowork/testdynamodb/myUpload"
)



//var uploadTemplate1 = template.Must(template.ParseFiles("login.html"))

func indexHandle1(w http.ResponseWriter, r *http.Request) {

	tpl ,err:=template.ParseFiles("login.html")
	if err = tpl.Execute(w, nil); err != nil {
		log.Fatal("Execute: ", err.Error())
		return
	}
}
func indexHandle2(w http.ResponseWriter, r *http.Request) {

	tpl ,err:=template.ParseFiles("index1.html")
	if err = tpl.Execute(w, nil); err != nil {
		log.Fatal("Execute: ", err.Error())
		return
	}
}
func indexHandle3(w http.ResponseWriter, r *http.Request) {

	tpl ,err:=template.ParseFiles("photolist.html")
	if err = tpl.Execute(w, nil); err != nil {
		log.Fatal("Execute: ", err.Error())
		return
	}
}



func main()  {
	//createUserTab()
	//PutUserItem("muser","mpassword")
	//scanUser()
	//getPasswordItem("muser")
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", indexHandle1)
	http.HandleFunc("/uploadpage", indexHandle2)
	http.HandleFunc("/querypage", indexHandle3)

	http.HandleFunc("/login", login)
	http.HandleFunc("/upload", myUpload.UploadHandle)
	http.HandleFunc("/queryall", myUpload.QueryAll)
	//http.ListenAndServe("10.105.192.98:8080", nil)
	http.ListenAndServe("10.105.192.98:8084", nil)
}
func queryallImage(){

		sess, err := session.NewSession(&aws.Config{
			Region:      aws.String("cn-northwest-1"),
			Credentials: credentials.NewStaticCredentials("AKIAPFDFDDUVOMJRDFTA",
				"bGhBzgVXQQfRgm+U06fVLu44WEVZVdL285wnyXPz", ""),
		})

		svc := dynamodb.New(sess)
		input := &dynamodb.ScanInput{
			//FilterExpression:     aws.String("Artist = :a"),
			ProjectionExpression: aws.String("imageurl"),
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

func login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*") //允许跨域
	r.ParseForm()
	username, found1 := r.Form["username"]
	password, found2 := r.Form["password"]

	if !(found1 && found2) {
		io.WriteString(w, "请勿非法访问")
		return
	}
	fmt.Println("password=",password)
	dbpassword :=getPasswordItem(username[0])
	fmt.Println("dbpassword=",dbpassword)

	if dbpassword == password[0] {
		http.Redirect(w, r, "/uploadpage", http.StatusFound)
	}else {
		fmt.Println("登陆失败")
	}

}




type Item struct {
	Username string       // Hash key, a.k.a. partition key
	Password  string // Range key, a.k.a. sort ke
}
func getPasswordItem(user string) string{
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("cn-northwest-1"),
		Credentials: credentials.NewStaticCredentials("AKIAPFDFDDUVOMJRDFTA",
			"bGhBzgVXQQfRgm+U06fVLu44WEVZVdL285wnyXPz", ""),
	})

	svc := dynamodb.New(sess)
	input := &dynamodb.QueryInput{
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":v1": {
				S: aws.String(user),
			},
		},
		KeyConditionExpression: aws.String("Username = :v1"),
		//ProjectionExpression:   aws.String("Password"),
		TableName:              aws.String("user"),
	}

	//result, err := svc.GetItem(input)
	result, err := svc.Query(input)

	items := []Item{}

	// Unmarshal the Items field in the result value to the Item Go type.
	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, &items)

	if err != nil {
		return ""
	}

	fmt.Println(items)
	return items[0].Password
}



func scanUser(){
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("cn-northwest-1"),
		Credentials: credentials.NewStaticCredentials("AKIAPFDFDDUVOMJRDFTA",
			"bGhBzgVXQQfRgm+U06fVLu44WEVZVdL285wnyXPz", ""),
	})

	svc := dynamodb.New(sess)
	input := &dynamodb.ScanInput{
		TableName:            aws.String("user"),
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




func createUserTab(){
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("cn-northwest-1"),
		Credentials: credentials.NewStaticCredentials("AKIAPFDFDDUVOMJRDFTA",
			"bGhBzgVXQQfRgm+U06fVLu44WEVZVdL285wnyXPz", ""),
	})

	svc := dynamodb.New(sess)
	input := &dynamodb.CreateTableInput{
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("Username"),
				AttributeType: aws.String("S"),
			},
			{
				AttributeName: aws.String("Password"),
				AttributeType: aws.String("S"),
			},
		},
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("Username"),
				KeyType:       aws.String("HASH"),
			},
			{
				AttributeName: aws.String("Password"),
				KeyType:       aws.String("RANGE"),
			},
		},
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(5),
			WriteCapacityUnits: aws.Int64(5),
		},
		TableName: aws.String("user"),
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

func PutUserItem(user ,password string) {
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("cn-northwest-1"),
		Credentials: credentials.NewStaticCredentials("AKIAPFDFDDUVOMJRDFTA",
			"bGhBzgVXQQfRgm+U06fVLu44WEVZVdL285wnyXPz", ""),
	})

	svc := dynamodb.New(sess)
	input := &dynamodb.PutItemInput{
		Item: map[string]*dynamodb.AttributeValue{
			"Username": {
				S: aws.String(user),
			},
			"Password": {
				S: aws.String(password),
			},

		},
		ReturnConsumedCapacity: aws.String("TOTAL"),
		TableName:              aws.String("user"),
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
