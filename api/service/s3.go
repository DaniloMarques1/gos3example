package service

import (
	"io"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type S3Service struct {
	session *session.Session
}

func NewS3() *S3Service {
	session := session.New(&aws.Config{
		Region:           aws.String(os.Getenv("AWS_REGION")),
		Endpoint:         aws.String(os.Getenv("AWS_ENDPOINT")),
		S3ForcePathStyle: aws.Bool(true),
	})

	return &S3Service{session}
}

func (s *S3Service) PutObject(name string, obj io.ReadSeeker) error {
	input := &s3.PutObjectInput{
		Bucket: aws.String(os.Getenv("S3_BUCKET_NAME")),
		Body:   obj,
		Key:    aws.String(name),
	}

	if _, err := s3.New(s.session).PutObject(input); err != nil {
		return err
	}

	return nil
}
