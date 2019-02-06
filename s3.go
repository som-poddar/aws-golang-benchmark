package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func ListBuckets() {
	svc := s3.New(session.New())
	input := &s3.ListBucketsInput{}

	_, err := svc.ListBuckets(input)
	if err != nil {
		fmt.Println(err.Error())
	}

	// fmt.Println(result)
}
