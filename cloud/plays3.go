package cloud

import (
	"fmt"
	"os"
	"sync"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

var filename = "./cloud/plays3.go"
var bucket = "saotomek-playground-bucket"
var key = "test/plays3.go"
var sess = session.Must(session.NewSession(aws.NewConfig().WithRegion("ap-northeast-1")))
var s3svc = s3.New(sess)
var uploader = s3manager.NewUploader(sess)

func PlayUploadingFile() {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	input := &s3manager.UploadInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
		Body:   file,
	}
	output, err := uploader.Upload(input)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("output: ", output)
	}
}

func DeleteFile() {
	request := &s3.DeleteObjectInput{
		Bucket: &bucket,
		Key:    &key,
	}
	output, err := s3svc.DeleteObject(request)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("output: ", output)
	}
}

func ObtainLock() {
	outputc := make(chan *s3.DeleteObjectOutput)
	errc := make(chan error)
	request := &s3.DeleteObjectInput{
		Bucket: &bucket,
		Key:    &key,
	}
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		output, err := s3svc.DeleteObject(request)
		if err != nil {
			errc <- err
		} else {
			outputc <- output
		}
		wg.Done()
	}()
	go func() {
		output, err := s3svc.DeleteObject(request)
		if err != nil {
			errc <- err
		} else {
			outputc <- output
		}
		wg.Done()
	}()
	go func() {
		wg.Wait()
		close(outputc)
		close(errc)
	}()
	for o := range outputc {
		fmt.Printf("marker: %v, charged: %v, version: %v\n", o.DeleteMarker, o.RequestCharged, o.VersionId)
	}
	for e := range errc {
		fmt.Println("error: ", e)
	}
}

func TagFile() {
	request1 := &s3.PutObjectTaggingInput{
		Bucket: &bucket,
		Key:    &key,
		Tagging: &s3.Tagging{
			TagSet: []*s3.Tag{
				{
					Key:   aws.String("hellios.UsedBy"),
					Value: aws.String("request1"),
				},
			},
		},
	}
	/**
	request2 := &s3.PutObjectTaggingInput{
		Bucket:  &bucket,
		Key:     &key,
		Tagging: &s3.Tagging{
			TagSet: []*s3.Tag {
				{
					Key: "hellios.UsedBy"
					Value: "request2"
				}
			}
		},
	}
	*/
	output, err := s3svc.PutObjectTagging(request1)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("output: ", output)
	}
}

func PlayLegalHold() {
	request := &s3.PutObjectLegalHoldInput{
		Bucket: &bucket,
		Key:    &key,
		LegalHold: &s3.ObjectLockLegalHold{
			Status: aws.String(s3.ObjectLockLegalHoldStatusOn),
		},
	}
	var wg sync.WaitGroup
	wg.Add(2)
	oc := make(chan *s3.PutObjectLegalHoldOutput, 2)
	ec := make(chan error, 2)
	for i := 0; i < 2; i++ {
		go func() {
			output, err := s3svc.PutObjectLegalHold(request)
			if err != nil {
				ec <- err
			} else {
				oc <- output
			}
			wg.Done()
		}()
	}
	fmt.Println("who?")
	go func() {
		wg.Wait()
		close(oc)
		close(ec)
	}()
	for e := range ec {
		fmt.Println("err: ", e)
	}
	for o := range oc {
		fmt.Println("output: ", o)
	}
}
