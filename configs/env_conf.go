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
	DbHost             string
	DbPort             string
	DbUser             string
	DbPassword         string
	DbName             string
)

func Init() error {
	err := godotenv.Load(".env")
	if err != nil {
		return err
	}
	DbHost = os.Getenv("DB_HOST")
	DbPort = os.Getenv("DB_PORT")
	DbUser = os.Getenv("DB_USER")
	DbPassword = os.Getenv("DB_PASSWORD")
	DbName = os.Getenv("DB_NAME")
	AwsAccessKeyId = os.Getenv("AWS_ACCESS_KEY_ID")
	AwsSecretAccessKey = os.Getenv("AWS_SECRET_ACCESS_KEY")
	AwsRegion = os.Getenv("AWS_REGION")
	BucketName = os.Getenv("BUCKET_NAME")

	return nil
}
