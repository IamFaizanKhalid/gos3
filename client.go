package gos3

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"io/ioutil"
	"net/http"
)

type s3Client struct {
	config   *Config
	handle   *s3.S3
	uploader *s3manager.Uploader
}

// NewClient returns a Client to use s3
func NewClient(config *Config) (Client, error) {
	sess, err := session.NewSession(&aws.Config{
		Region:           aws.String(config.Region),
		Endpoint:         aws.String(config.Endpoint),
		S3ForcePathStyle: aws.Bool(true),
		Credentials: credentials.NewStaticCredentials(
			config.AccessKeyId,
			config.SecretAccessKey,
			"",
		),
	})
	if err != nil {
		return nil, err
	}

	return &s3Client{
		config:   config,
		handle:   s3.New(sess),
		uploader: s3manager.NewUploader(sess),
	}, nil
}

func (client *s3Client) CheckConnection() error {
	res, err := http.DefaultClient.Get(client.config.Endpoint)
	if err != nil {
		return fmt.Errorf("Cannot connect to s3..")
	}
	defer res.Body.Close()

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	if string(b) != "{\"status\": \"running\"}" {
		return fmt.Errorf("%s", string(b))
	}

	return nil
}

func (client *s3Client) SelectBucket(bucketName string) (Bucket, error) {
	if client.CheckConnection() != nil {
		return nil, fmt.Errorf("Cannot connect to s3..")
	}

	return &s3Bucket{
		client:     client,
		bucketName: bucketName,
	}, nil
}
