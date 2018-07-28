package main


/*
	https://www.youtube.com/watch?v=iOGIKG3EptI
	https://github.com/awslabs/aws-go-wordfreq-sample/blob/master/cmd/uploads3/main.go

	https://docs.aws.amazon.com/sdk-for-go/api/aws/
	- first configure your aws credentials run: aws configure
	- go get -u github.com/aws/aws-sdk-go/aws
	- login to UI web aws s3 interface
	- go to S3 service
	- create a Bucket called com.example in the desired region (I used Oregon us-west-2)
	- run:   go run main.go com.example fileToUpload
*/


import (
"fmt"
"os"
"github.com/aws/aws-sdk-go/aws"
"github.com/aws/aws-sdk-go/aws/session"
"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/aws/aws-sdk-go/aws/credentials"
)

func main() {
	//if len(os.Args) != 3 {
	//	fmt.Printf("usage: %s <bucket> <filename>\n", filepath.Base(os.Args[0]))
	//	os.Exit(1)
	//}

	bucket := "mubucket"
	filename := "C:\\Users\\green\\Desktop\\muwenbo\\white6.png"

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Failed to open file", filename, err)
		os.Exit(1)
	}
	defer file.Close()

	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("cn-northwest-1"),
		Credentials: credentials.NewStaticCredentials("AKIAPFDFDDUVOMJRDFTA",
			"bGhBzgVXQQfRgm+U06fVLu44WEVZVdL285wnyXPz", ""),
	})


	svc := s3manager.NewUploader(sess)

	fmt.Println("Uploading file to S3...")
	result, err := svc.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucket),
		ACL:    				aws.String("public-read"),
		Key:    aws.String("white7.png"),
		Body:   file,
	})
	if err != nil {
		fmt.Println("error", err)
		os.Exit(1)
	}

	fmt.Printf("Successfully uploaded %s to %s\n", filename, result.Location)
}
