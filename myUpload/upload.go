package myUpload




import (
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"bytes"
	"fmt"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

var uploadTemplate = template.Must(template.ParseFiles("login.html"))

func indexHandle(w http.ResponseWriter, r *http.Request) {
	if err := uploadTemplate.Execute(w, nil); err != nil {
		log.Fatal("Execute: ", err.Error())
		return
	}
}

func UploadHandle(w http.ResponseWriter, r *http.Request) {
	file, hander, err := r.FormFile("file")
	fmt.Println(hander.Filename)
	if err != nil {
		log.Fatal("FormFile: ", err.Error())
		return
	}
	defer func() {
		if err := file.Close(); err != nil {
			log.Fatal("Close: ", err.Error())
			return
		}
	}()
	//bytes, err := ioutil.ReadAll(file)
	//if err != nil {
	//	log.Fatal("ReadAll: ", err.Error())
	//	return
	//}

	defer file.Close()
	//fileInfo, _ := file.Stat()
	//var size = fileInfo.Size()

	//path := "/media/" + file.Name()
	buffer, err := ioutil.ReadAll(file)
	file.Read(buffer)
	fileBytes := bytes.NewReader(buffer)
	fileType := http.DetectContentType(buffer)
	fmt.Println(fileType)
	input := s3.PutObjectInput{
		Bucket: aws.String("mubucket"),
		Key:    aws.String(hander.Filename),
		ACL:    aws.String("public-read"),
		Body:   fileBytes,
		//ContentLength: aws.Int64(size),
		ContentType: aws.String(fileType),
	}
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("cn-northwest-1"),
		Credentials: credentials.NewStaticCredentials("AKIAPFDFDDUVOMJRDFTA",
			"bGhBzgVXQQfRgm+U06fVLu44WEVZVdL285wnyXPz", ""),
	})

	svc := s3.New(sess)

	result, err := svc.PutObject(&input)

	//result, err := svc.CopyObject(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case s3.ErrCodeObjectNotInActiveTierError:
				fmt.Println(s3.ErrCodeObjectNotInActiveTierError, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		//return
	}

	fmt.Println(result)
	url:="https://mubucket.s3.cn-northwest-1.amazonaws.com.cn/"+hander.Filename
	svcdb := dynamodb.New(sess)
	inputdb := &dynamodb.PutItemInput{
		Item: map[string]*dynamodb.AttributeValue{
			"Imageurl": {
				S: aws.String(url),
			},
			"ImageName": {
				S: aws.String(hander.Filename),
			},
		},
		ReturnConsumedCapacity: aws.String("TOTAL"),
		TableName:              aws.String("MuImage"),
	}

	resultdb, err := svcdb.PutItem(inputdb)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(resultdb)


	w.Write([]byte(" <html> <body> <a> https://mubucket.s3.cn-northwest-1.amazonaws.com.cn/"+hander.Filename+" </a>  </body> </html>"))
	w.Write([]byte(" 已经存入数据库"))
}

//func main() {
//	http.HandleFunc("/", indexHandle)
//	//http.HandleFunc("/upload", uploadHandle)
//	http.ListenAndServe("10.105.192.98:8083", nil)
//}
