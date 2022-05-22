package controller

import (
	"net/http"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/faridlan/golang-upload-do/helper"
)

func UploadImage(writer http.ResponseWriter, request *http.Request) {

	s3Client := helper.S3Config()
	x, y := helper.Image(writer, request)
	z := "anjing"
	a := "-goblog1"
	az := z + a + ".png"

	object := s3.PutObjectInput{
		Bucket: aws.String("olshop"),
		Key:    aws.String("/profiles/" + az),
		Body:   strings.NewReader(string(y)),
		ACL:    aws.String("public-read"),
		Metadata: map[string]*string{
			"x-amz-acl":   aws.String("public-read"),
			"uploaded-by": aws.String("Faridlan"),
			"size":        aws.String(x),
		},
	}

	_, err := s3Client.PutObject(&object)
	if err != nil {
		panic(err)
	}
}
