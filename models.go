package gos3

import (
	"io"
	"mime/multipart"
	"time"
)

type Client interface {
	CheckConnection() error
	SelectBucket(bucketName string) (Bucket, error)
}

type Bucket interface {
	Upload(file multipart.File, fileName string, destination string) error
	List(directory string) ([]File, error)
	Get(filePath string) (io.Reader, error)
	Delete(filePath string) error
	GetPreSignedLink(filePath string, expiry time.Duration) (string, error)
}

type Config struct {
	Endpoint        string
	Region          string
	AccessKeyId     string
	SecretAccessKey string
}

type File struct {
	Name         string    `json:"name"`
	Size         int64     `json:"size"`
	LastModified time.Time `json:"last_modified"`
}
