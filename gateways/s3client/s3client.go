package s3client

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type s3gateway struct {
	config Config
	client *s3.Client
}

type Config struct {
	AWSRegion string
	AWSBucket string
}

// nolint
func New(config Config, s3Client *s3.Client) *s3gateway {
	return &s3gateway{
		config: config,
		client: s3Client,
	}
}

func (c *s3gateway) UploadToS3(filename string, key string) (*manager.UploadOutput, error) {
	uploader := manager.NewUploader(c.client)

	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Failed to open file %v", err)
		return nil, err
	}

	// upload to s3
	result, err := uploader.Upload(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(c.config.AWSBucket),
		Key:    aws.String(key),
		Body:   file,
	})
	if err != nil {
		fmt.Printf("Error uploading file: %v\n", err)
		return nil, err
	}

	return result, nil
}
