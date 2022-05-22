package controller

import (
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/faridlan/golang-upload-do/helper"
)

func DeleteImage(writer http.ResponseWriter, request *http.Request) {
	s3Client := helper.S3Config()
	input := &s3.DeleteObjectInput{
		Bucket: aws.String("olshop"),
		Key:    aws.String("/products/Anjay1.png"),
	}

	_, err := s3Client.DeleteObject(input)
	if err != nil {
		panic(err)
	}
}
