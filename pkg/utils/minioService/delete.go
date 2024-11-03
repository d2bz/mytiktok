package minioService

import (
	"context"
	"errors"
	"net/url"
	"strings"

	"github.com/minio/minio-go/v7"
)

func DeleteFileByURL(fileURL string) error {
	u, err := url.Parse(fileURL)
	if err != nil {
		return err
	}
	//从 URL 的路径中去掉前导的 /，然后用 / 分割，得到 bucketName 和 objectName。
	pathSegments := strings.Split(strings.TrimPrefix(u.Path, "/"), "/")
	if len(pathSegments) < 2 {
		return errors.New("URL 格式不正确，无法提取 bucketName 和 objectName")
	}
	// 提取 bucketName 和 objectName
	bucketName := pathSegments[0]
	objectName := strings.Join(pathSegments[1:], "/") // 处理包含 `/` 的 objectName

	err = minioClient.RemoveObject(context.Background(), bucketName, objectName, minio.RemoveObjectOptions{})
	if err != nil {
		return err
	}

	return nil
}
