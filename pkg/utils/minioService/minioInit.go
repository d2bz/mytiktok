package minioService

import (
	"context"
	"log"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

const (
	minioEndpoint   = "124.71.229.101:9000"
	accessKeyID     = "access_key_perfric"
	secretAccessKey = "secret_key_perfric"
	bucketName      = "mybucket"
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
