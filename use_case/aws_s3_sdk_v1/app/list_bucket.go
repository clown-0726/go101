package app

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"os"
)

var (
	sess  *session.Session
	svcS3 *s3.S3
)

func init() {
	accessKey := ""
	secretKey := ""

	// Create s3 session
	sess, _ = session.NewSession(&aws.Config{
		Credentials:      credentials.NewStaticCredentials(accessKey, secretKey, ""),
		Region:           aws.String("us-west-2"),
		DisableSSL:       aws.Bool(true),
		S3ForcePathStyle: aws.Bool(false), // virtual-host style 方式，不要修改
	})

	svcS3 = s3.New(sess)
}

func Main4ListBucket() {
	ListBuckets()
}

func ListBuckets() {
	// List all buckets
	result, err := svcS3.ListBuckets(nil)
	if err != nil {
		exitErrorf("Unable to list buckets")
	}

	fmt.Println("Now to list bucket...")

	for _, b := range result.Buckets {
		fmt.Println(aws.StringValue(b.Name))
	}
}

func exitErrorf(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, msg+"\n", args...)
	os.Exit(1)
}
