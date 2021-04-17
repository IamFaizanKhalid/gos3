# gos3 [![Build Status](https://api.travis-ci.com/IamFaizanKhalid/gos3.svg?branch=master)](https://travis-ci.com/github/IamFaizanKhalid/gos3) [![Go Report Card](https://goreportcard.com/badge/github.com/IamFaizanKhalid/gos3)](https://goreportcard.com/report/github.com/IamFaizanKhalid/gos3) ![License](https://img.shields.io/badge/license-MIT-blue.svg)
<img align="right" src="https://1.bp.blogspot.com/-AexS3vrfv_U/XSDuQks44aI/AAAAAAAAIsk/tuXxd6Cbc5Y6b8FjNFj8mQg0Oj7W1QQJQCLcBGAs/s1600/1_4lGHHzMSjWBrAHj1VGuXKQ.png" width="250">

Using S3 made easier.
Use S3 as you use your local file system with this simple wrapper over `github.com/aws/aws-sdk-go`.

## Includes
- Uploading a file
- Deleting a file
- Getting a pre-signed link to the file

## Example
```go
package main

import (
    "github.com/IamFaizanKhalid/gos3"
    "log"
    "os"
    "time"
)

func main() {
    // Getting client
    client, err := gos3.NewClient(&gos3.Config{
        Endpoint:        "http://localhost:4566",
        Region:          "us-east-1",
        AccessKeyId:     "test",
        SecretAccessKey: "test",
    })
    if err != nil {
        log.Panicln(err)
    }

    // Select a bucket    
    bucket, err := client.SelectBucket("test-bucket")
    if err != nil {
        log.Panicln(err)
    }

    // Upload a file
    file, err := os.Open("README.md")
    if err != nil {
        log.Panicln(err)
    }

    err = bucket.Upload(file, file.Name(), "backup")
    if err != nil {
        log.Panicln(err)
    }

    // Get pre-signed link to the file
    link, err := bucket.GetPreSignedLink("backup/README.md", time.Hour)
    if err != nil {
        log.Panicln(err)
    }

    log.Printf("Pre-signed link: %s\n", link)

    // Delete the file
    err = bucket.Delete("backup/README.md")
    if err != nil {
        log.Panicln(err)
    }
}
```
