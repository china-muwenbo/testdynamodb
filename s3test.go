package main



import (

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"fmt"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"os"
	"bytes"
	"net/http"
)

func main()  {

	s3GetAll()

}

func s3GetAll(){
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("cn-northwest-1"),
		Credentials: credentials.NewStaticCredentials("AKIAPFDFDDUVOMJRDFTA",
			"bGhBzgVXQQfRgm+U06fVLu44WEVZVdL285wnyXPz", ""),
	})

	svc := s3.New(sess)
	//input := &s3.CopyObjectInput{
	//
	//	Bucket:     aws.String("mubucket"),
	//	CopySource: aws.String("/mubucket/white.png"),
	//	Key:        aws.String("white.png"),
	//}

	//path := "C:\\Users\\green\\Desktop\\muwenbo\\white1.png"
	//fd, _ := os.OpenFile(path, os.O_RDWR|os.O_APPEND, 0666)
	//////rd := bufio.NewReaderSize(fd, 4096)
	////

	file, err := os.Open("C:\\Users\\green\\Desktop\\muwenbo\\black.png")
	if err != nil {
		// handle error
		fmt.Println(err)
	}
	defer file.Close()
	fileInfo, _ := file.Stat()
	var size = fileInfo.Size()

	//path := "/media/" + file.Name()
	buffer := make([]byte, size)
	file.Read(buffer)
	fileBytes := bytes.NewReader(buffer)
	fileType := http.DetectContentType(buffer)
	fmt.Println(fileType)
	input := &s3.PutObjectInput{
		Bucket: aws.String("mubucket"),
		Key: aws.String("white9.png"),
		ACL:    				aws.String("public-read"),
		Body: fileBytes,
		ContentLength: aws.Int64(size),
		ContentType: aws.String(fileType),
	}
	//input := &s3.PutObjectInput{
	//	Body:                 fd,
	//	Bucket:               aws.String("mubucket"),
	//	Key:                  aws.String("white4.png"),
	//
	//	GrantRead:           aws.String("uri=http://acs.amazonaws.com/groups/global/AllUsers"),
	//	//ServerSideEncryption: aws.String("AES256"),
	//	//Tagging:              aws.String("key1=value1&key2=value2"),
	//
	//}

	//input := &s3.PutObjectInput{
	//	Body:    aws.ReadSeekCloser(strings.NewReader("C:\\Users\\green\\Desktop\\muwenbo\\white1.png")),
	//	Bucket:  aws.String("mubucket"),
	//	Key:     aws.String("white5.png"),
	//	Tagging: aws.String("key1=value1&key2=value2"),
	//}
	//svc.
	result, err := svc.PutObject(input)

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
		return
	}

	fmt.Println(result)
}

func s3Upload(){

}

func s3GetItem(){
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("cn-northwest-1"),
		Credentials: credentials.NewStaticCredentials("AKIAPFDFDDUVOMJRDFTA",
			"bGhBzgVXQQfRgm+U06fVLu44WEVZVdL285wnyXPz", ""),
	})

	svc := s3.New(sess)

	input1 := &s3.GetObjectInput{
		Bucket: aws.String("mubucket"),
		Key:    aws.String("dlrb.png"),
	}

	result1, err := svc.GetObject(input1)

	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case s3.ErrCodeNoSuchKey:
				fmt.Println(s3.ErrCodeNoSuchKey, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
	}

	fmt.Println(result1)
}