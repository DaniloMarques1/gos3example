package main

import (
	"context"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	s3Session := session.New(&aws.Config{
		Region: aws.String(os.Getenv("AWS_REGION")),
		Endpoint: aws.String(os.Getenv("AWS_ENDPOINT")),
		S3ForcePathStyle: aws.Bool(true),
	})

	uploader := s3manager.NewUploader(s3Session)

	fileName := "cat.jpg"
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}

	input := &s3manager.UploadInput{
		Bucket: aws.String(os.Getenv("S3_BUCKET_NAME")),
		Body: f,
		Key: aws.String(fileName),
	}

	if _, err := uploader.UploadWithContext(context.Background(), input); err != nil {
		log.Fatal(err)
	}
}

