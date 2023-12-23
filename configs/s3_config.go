package configs

import (
	"github.com/joho/godotenv"
	"os"
)

var (
	AwsAccessKeyId     string
	AwsSecretAccessKey string
	AwsRegion          string
	BucketName         string
)

func Init() error {
	err := godotenv.Load(".env")
	if err != nil {
		return err
	}

	AwsAccessKeyId = os.Getenv("AWS_ACCESS_KEY_ID")
	AwsSecretAccessKey = os.Getenv("AWS_SECRET_ACCESS_KEY")
	AwsRegion = os.Getenv("AWS_REGION")
	BucketName = os.Getenv("BUCKET_NAME")

	return nil
}
