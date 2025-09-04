package s3

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
	"io"
	"path/filepath"
	"strings"
	"time"
)

// UploadStream 上传流（适合 Web 上传的文件）
func UploadStream(ctx context.Context, bucketName, objectName string, reader io.Reader, objectSize int64, contentType string) (minio.UploadInfo, error) {
	c := GetClient()
	info, err := c.PutObject(ctx, bucketName, GenerateObjectName(objectName), reader, objectSize, minio.PutObjectOptions{
		ContentType: contentType,
	})
	if err != nil {
		return minio.UploadInfo{}, err
	}
	return info, nil
}

// DownloadStream 获取对象的流（适合直接返回给 HTTP 响应）
func DownloadStream(ctx context.Context, bucketName, objectName string) (io.ReadCloser, error) {
	c := GetClient()
	return c.GetObject(ctx, bucketName, objectName, minio.GetObjectOptions{})
}

// GenerateObjectName 生成唯一的对象名称，包含日期路径和 UUID，防止命名冲突且后续可以按日期归档清理
func GenerateObjectName(filename string) string {
	now := time.Now()
	datePath := now.Format("200601")
	unique := uuid.New().String()[:8]
	return fmt.Sprintf("%s/%s-%s", datePath, unique, filepath.Base(filename))
}

func GetFileExt(filename string) string {
	parts := strings.Split(filename, ".")
	if len(parts) > 1 {
		return parts[len(parts)-1]
	}
	return ""
}
