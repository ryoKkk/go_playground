package play

import (
	"fmt"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func DownloadFromS3() {
	sess := session.Must(session.NewSession(aws.NewConfig().WithRegion("ap-northeast-1")))
	downloader := s3manager.NewDownloader(sess)
	key := "paper/amazon-dynamo-sosp2007.pdf"
	f, err := os.Create(getFilenameOnly(key))
	if err != nil {
		panic(err)
	}
	n, err := downloader.Download(f, &s3.GetObjectInput{
		Bucket: aws.String("saotomek-bucket"),
		Key:    aws.String(key),
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("downloaded bytes: ", n)
}

func getFilenameOnly(path string) string {
	eles := strings.Split(path, "/")
	return eles[len(eles)-1]
}

func PathTest() {
	fmt.Println("file: ", getFilenameOnly("a/b/c.txt"))
	fmt.Println("file: ", getFilenameOnly("b/c.txt"))
	fmt.Println("file: ", getFilenameOnly("c.txt"))
}
