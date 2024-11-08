package minioService

import (
	"context"
	"log"
	"tiktok/config"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

const (
	minioEndpoint   = config.MINIO_ADDR
	accessKeyID     = config.ACCESS_KEY_ID
	secretAccessKey = config.SECRET_ACCESS_KEY
	bucketName      = config.BUCKET_NAME
)

var minioClient *minio.Client
var cbg = context.Background()

func MinioInit() {
	//初始化客户端
	var err error
	minioClient, err = minio.New(minioEndpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: false,
	})
	if err != nil {
		log.Fatalln("minio初始化客户端错误：" + err.Error())
	}

	//确保bucket存在
	err = minioClient.MakeBucket(cbg, bucketName, minio.MakeBucketOptions{Region: "us-east-1"})
	if err != nil {
		exists, errBucketExists := minioClient.BucketExists(cbg, bucketName)
		if errBucketExists == nil && exists {
			log.Printf("Bucket %s already exists\n", bucketName)
		} else {
			log.Fatalln("minio创建bucket错误：" + err.Error())
		}
	} else {
		log.Printf("Bucket %s created successfully\n", bucketName)
	}
}
