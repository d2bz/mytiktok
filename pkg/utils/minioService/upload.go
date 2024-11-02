package minioService

import (
	"context"
	"errors"
	"fmt"
	"mime/multipart"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
)

func uploadFile(file *multipart.FileHeader) (string, error) {
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

func UploadVideo(c *gin.Context) (string, error) {
	// 从请求中获取文件
	videoFile, err := c.FormFile("video")
	if err != nil {
		return "", errors.New("file is required")
	}

	//格式判断
	//".mp4", ".avi", ".mov", ".mkv"
	ext := filepath.Ext(videoFile.Filename)
	if ext != ".mp4" && ext != ".avi" && ext != ".mov" && ext != ".mkv" {
		return "", errors.New("视频格式错误")
	}

	url, err := uploadFile(videoFile)
	if err != nil {
		return "", err
	}

	return url, nil
}
