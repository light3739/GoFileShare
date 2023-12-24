package storage

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/google/uuid"
	"mime/multipart"
	"net/http"
	"share-files/configs"
)

func NewAWSSession() (*session.Session, error) {
	return session.NewSession(&aws.Config{
		Region:      aws.String(configs.AwsRegion),
		Credentials: credentials.NewStaticCredentials(configs.AwsAccessKeyId, configs.AwsSecretAccessKey, ""),
	})
}

func GetFileFromRequest(r *http.Request) (multipart.File, *multipart.FileHeader, error) {
	return r.FormFile("file")
}

func UploadFile(sess *session.Session, file multipart.File, header *multipart.FileHeader) error {
	uploader := s3manager.NewUploader(sess)
	_, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(configs.BucketName),
		Key:    aws.String(header.Filename),
		Body:   file,
	})
	return err
}

func GenerateTag() string {
	return uuid.New().String()
}

func AddTag(sess *session.Session, filename string, tag string) error {
	svc := s3.New(sess)
	_, err := svc.PutObjectTagging(&s3.PutObjectTaggingInput{
		Bucket: aws.String(configs.BucketName),
		Key:    aws.String(filename),
		Tagging: &s3.Tagging{
			TagSet: []*s3.Tag{
				{
					Key:   aws.String("uniqueTag"),
					Value: aws.String(tag),
				},
			},
		},
	})
	return err
}
