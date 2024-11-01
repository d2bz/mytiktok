package utils

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
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
		log.Fatalln(err)
	}

	//确保bucket存在
	err = minioClient.MakeBucket(cbg, bucketName, minio.MakeBucketOptions{Region: "us-east-1"})
	if err != nil {
		exists, errBucketExists := minioClient.BucketExists(cbg, bucketName)
		if errBucketExists == nil && exists {
			log.Printf("Bucket %s already exists\n", bucketName)
		} else {
			log.Fatalln(err)
		}
	}
}

func UploadFile(c *gin.Context) (string, error) {
	// 从请求中获取文件
	file, err := c.FormFile("file")
	if err != nil {
		return "", errors.New("file is required")
	}

	// 打开文件
	src, err := file.Open() //src是一个实现了 io.Reader 接口的对象，可以通过它读取文件数据。
	if err != nil {
		return "", errors.New("unable to open file")
	}
	defer src.Close()

	// 上传到 MinIO
	objectName := file.Filename
	contentType := file.Header.Get("Content-Type") //得到文件的类型
	_, err = minioClient.PutObject(context.Background(), bucketName, objectName, src, file.Size, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		return "", errors.New("unable to upload file")
	}

	// 构建文件的 URL
	fileURL := fmt.Sprintf("http://%s/%s/%s", minioEndpoint, bucketName, objectName)
	return fileURL, nil
}
