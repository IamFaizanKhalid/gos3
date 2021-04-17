package gos3

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"io"
	"mime/multipart"
	"time"
)

type s3Bucket struct {
	client     *s3Client
	bucketName string
}

func (bucket *s3Bucket) Upload(file multipart.File, fileName string, destination string) error {
	_, err := bucket.client.uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucket.bucketName),
		ACL:    aws.String("private"),
		Key:    aws.String(destination + "/" + fileName),
		Body:   file,
	})

	return err
}

func (bucket *s3Bucket) List(directory string) ([]File, error) {
	resp, err := bucket.client.handle.ListObjectsV2(&s3.ListObjectsV2Input{
		Bucket: aws.String(bucket.bucketName),
		Prefix: aws.String(directory),
	})
	if err != nil {
		return nil, err
	}

	var files []File
	for _, item := range resp.Contents {
		files = append(files, File{
			Name:         *item.Key,
			Size:         *item.Size,
			LastModified: *item.LastModified,
		})
	}

	return files, nil
}

func (bucket *s3Bucket) Get(filePath string) (io.Reader, error) {
	req, out := bucket.client.handle.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String(bucket.bucketName),
		Key:    aws.String(filePath),
	})

	err := req.Send()
	if err != nil {
		return nil, err
	}

	return out.Body, nil
}

func (bucket *s3Bucket) GetPreSignedLink(filePath string, expiry time.Duration) (string, error) {
	req, _ := bucket.client.handle.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String(bucket.bucketName),
		Key:    aws.String(filePath),
	})

	return req.Presign(expiry)
}

func (bucket *s3Bucket) Delete(filePath string) error {
	_, err := bucket.client.handle.DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String(bucket.bucketName),
		Key:    aws.String(filePath),
	})

	return err
}
