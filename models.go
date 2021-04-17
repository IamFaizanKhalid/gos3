package gos3

import (
	"io"
	"mime/multipart"
	"time"
)

// Config used to connect to a s3 bucket
type Config struct {
	Endpoint        string
	Region          string
	AccessKeyId     string
	SecretAccessKey string
}

// Client represents a gos3 client
type Client interface {
	// CheckConnection returns nil when s3 can be accessible on the given endpoint,
	// or error otherwise
	CheckConnection() error

	// SelectBucket returns an object with a bucket selected.
	// It returns an error if connection to s3 fails.
	SelectBucket(bucketName string) (Bucket, error)
}

// Bucket represents a selected bucket from a client
type Bucket interface {
	// Upload a file to selected bucket
	Upload(file multipart.File, fileName string, destination string) error

	// List files in the given directory of the selected bucket
	List(directory string) ([]File, error)

	// Get file from the selected bucket
	Get(filePath string) (io.Reader, error)

	// Delete file from the selected bucket
	Delete(filePath string) error

	// GetPreSignedLink of a file from the selected bucket with an expiry time
	GetPreSignedLink(filePath string, expiry time.Duration) (string, error)
}

// File contains file information returned from List()
type File struct {
	Name         string    `json:"name"`
	Size         int64     `json:"size"`
	LastModified time.Time `json:"last_modified"`
}
