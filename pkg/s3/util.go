package s3

import (
	"context"
	"io"
	"time"

	"github.com/minio/minio-go/v7"
)

// UploadFile 上传本地文件到 S3
func UploadFile(ctx context.Context, bucketName, objectName, filePath, contentType string) (minio.UploadInfo, error) {
	c := GetClient()
	info, err := c.FPutObject(ctx, bucketName, objectName, filePath, minio.PutObjectOptions{
		ContentType: contentType,
	})
	if err != nil {
		return minio.UploadInfo{}, err
	}
	return info, nil
}

// UploadStream 上传流（适合 Web 上传的文件）
func UploadStream(ctx context.Context, bucketName, objectName string, reader io.Reader, objectSize int64, contentType string) (minio.UploadInfo, error) {
	c := GetClient()
	info, err := c.PutObject(ctx, bucketName, objectName, reader, objectSize, minio.PutObjectOptions{
		ContentType: contentType,
	})
	if err != nil {
		return minio.UploadInfo{}, err
	}
	return info, nil
}

// DownloadFile 下载对象到本地文件
func DownloadFile(ctx context.Context, bucketName, objectName, localPath string) error {
	c := GetClient()
	return c.FGetObject(ctx, bucketName, objectName, localPath, minio.GetObjectOptions{})
}

// DownloadStream 获取对象的流（适合直接返回给 HTTP 响应）
func DownloadStream(ctx context.Context, bucketName, objectName string) (io.ReadCloser, error) {
	c := GetClient()
	return c.GetObject(ctx, bucketName, objectName, minio.GetObjectOptions{})
}

// GetPresignedURL 生成临时访问链接
func GetPresignedURL(ctx context.Context, bucketName, objectName string, expiry time.Duration) (string, error) {
	c := GetClient()
	url, err := c.PresignedGetObject(ctx, bucketName, objectName, expiry, nil)
	if err != nil {
		return "", err
	}
	return url.String(), nil
}
