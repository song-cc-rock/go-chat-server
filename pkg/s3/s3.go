package s3

import (
	"context"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"go-chat-server/pkg/config"
	"log"
	"sync"
)

var (
	client *minio.Client
	once   sync.Once
)

// Init 初始化s3客户端
func Init() error {
	var err error
	once.Do(func() {
		endpoint := config.GetString("minio.endpoint")
		accessKeyID := config.GetString("minio.access_key")
		secretAccessKey := config.GetString("minio.secret_key")
		useSSL := config.GetBool("minio.ssl")

		client, err = minio.New(endpoint, &minio.Options{
			Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
			Secure: useSSL,
		})
		if err != nil {
			return
		}
		log.Println("✅ MinIO client connected:", endpoint)
	})
	return err
}

// GetClient 获取单例客户端
func GetClient() *minio.Client {
	if client == nil {
		panic("❌ MinIO client is not initialized, call s3.Init() first")
	}
	return client
}

// CreateBucket 创建桶（示例方法）
func CreateBucket(ctx context.Context, bucketName string) error {
	c := GetClient()

	exists, err := c.BucketExists(ctx, bucketName)
	if err != nil {
		return err
	}
	if !exists {
		err = c.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{})
		if err != nil {
			return err
		}
		log.Println("✅ Created bucket:", bucketName)
	} else {
		log.Println("⚠️ Bucket already exists:", bucketName)
	}
	return nil
}
