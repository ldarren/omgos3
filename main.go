package main
import (
	"fmt"
	"bytes"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/aws/credentials"
)

func main(){
	bucket := aws.String("kloudkonsole-log")
	keyname := aws.String("log.json")

	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"),
		Credentials: credentials.NewSharedCredentials("", "logger"),
	}))

	svc := s3.New(sess, &aws.Config{
		Region: aws.String("us-east-1")
	})

	params := &s3.PutObjectInput{
		Bucket: bucket,  // Required
		Key:    keyname, // Required
		ACL:    aws.String("bucket-owner-full-control"),
		Body:   bytes.NewReader([]byte("PAYLOAD")),
	// CacheControl:         aws.String("CacheControl"),
	// ContentDisposition:   aws.String("ContentDisposition"),
	// ContentEncoding:      aws.String("ContentEncoding"),
	// ContentLanguage:      aws.String("ContentLanguage"),
		ContentLength: aws.Int64(7),
	// ContentType:          aws.String("ContentType"),
	// Expires:              aws.Time(time.Now()),
	// GrantFullControl:     aws.String("GrantFullControl"),
	// GrantRead:            aws.String("GrantRead"),
	// GrantReadACP:         aws.String("GrantReadACP"),
	// GrantWriteACP:        aws.String("GrantWriteACP"),
	// RequestPayer:         aws.String("RequestPayer"),
	// SSECustomerAlgorithm: aws.String("SSECustomerAlgorithm"),
	// SSECustomerKey:       aws.String("SSECustomerKey"),
	// SSECustomerKeyMD5:    aws.String("SSECustomerKeyMD5"),
	// SSEKMSKeyId:          aws.String("SSEKMSKeyId"),
	}
	resp, err := svc.PutObject(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	fmt.Println(resp)
}
