package utils

import (
	"efeasy-gin/global"
	"fmt"
	"github.com/minio/minio-go"
	"github.com/minio/minio-go/pkg/policy"
	"go.uber.org/zap"
	"io"
	"net/url"
	"time"
)

// CreateMinoBucket 创建minio 桶
func CreateMinoBucket(bucketName string) {
	location := "us-east-1"
	err := global.App.MinioClient.MakeBucket(bucketName, location)
	if err != nil {
		// 检查存储桶是否已经存在。
		exists, err := global.App.MinioClient.BucketExists(bucketName)
		fmt.Println(exists)
		if err == nil && exists {
			fmt.Printf("We already own %s\n", bucketName)
		} else {
			fmt.Println(err, exists)
			return
		}
	}
	//
	err = global.App.MinioClient.SetBucketPolicy(bucketName, policy.BucketPolicyReadWrite)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Successfully created %s\n", bucketName)
}

// UploadFile 上传文件给minio指定的桶中
func UploadFile(bucketName, objectName string, reader io.Reader, objectSize int64) (ok bool) {
	n, err := global.App.MinioClient.PutObject(bucketName, objectName, reader, objectSize, minio.PutObjectOptions{ContentType: "application/octet-stream"})
	if err != nil {
		fmt.Println(err)
		return false
	}
	fmt.Println("Successfully uploaded bytes: ", n)
	return true
}

// GetFileUrl  获取文件url
func GetFileUrl(bucketName string, fileName string, expires time.Duration) string {
	//time.Second*24*60*60
	reqParams := make(url.Values)
	preSignedURL, err := global.App.MinioClient.PresignedGetObject(bucketName, fileName, expires, reqParams)
	if err != nil {
		zap.L().Error(err.Error())
		return ""
	}
	return fmt.Sprintf("%s", preSignedURL)
}
