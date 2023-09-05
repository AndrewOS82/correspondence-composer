package usecases

import (
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
)

type Uploader struct {
	StorageClient storageClientGateway
}

type storageClientGateway interface {
	UploadToS3(filename string, key string) (*manager.UploadOutput, error)
}

func (u *Uploader) Upload(filepath string, key string) error {
	_, err := u.StorageClient.UploadToS3(filepath, key)
	if err != nil {
		return err
	}

	return nil
}
