package initialize

import (
	"efeasy-gin/global"
	"efeasy-gin/utils"
	"github.com/minio/minio-go"
	_ "github.com/minio/minio-go/pkg/encrypt"
	"log"
)

func MinIO() {
	minioInfo := global.App.Config.MinIo
	// 初使化 minio client对象。false是关闭https证书校验
	minioClient, err := minio.New(minioInfo.Endpoint, minioInfo.AccessKeyId, minioInfo.SecretAccessKey, false)
	if err != nil {
		log.Fatalln(err)
	}
	//客户端注册到全局变量中
	global.App.MinioClient = minioClient
	//创建一个叫user header的存储桶。
	utils.CreateMinoBucket("user_header")
}
