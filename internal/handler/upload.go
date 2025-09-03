package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	v1 "go-chat-server/api/v1"
	"go-chat-server/internal/service"
	"go-chat-server/pkg/s3"
	"net/http"
)

type UploadHandler struct {
	uploadService service.UploadService
}

func NewUploadHandler() *UploadHandler {
	return &UploadHandler{
		uploadService: service.NewUploadService(),
	}
}

// Upload 上传文件
func (c *UploadHandler) Upload(ctx *gin.Context) {
	// 从表单获取文件
	file, err := ctx.FormFile("file")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "file not found"})
		return
	}

	// 打开文件流
	f, err := file.Open()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "cannot open file"})
		return
	}
	defer f.Close()

	// 上传到 S3
	bucket := "go-chat"
	objectName := file.Filename
	info, err := s3.UploadStream(context.Background(), bucket, objectName, f, file.Size, file.Header.Get("Content-Type"))
	if err != nil {
		v1.HandleError(ctx, http.StatusInternalServerError, "upload failed: "+err.Error())
		return
	}

	v1.HandleSuccess(ctx, gin.H{
		"fileName": info.Key,
		"size":     info.Size,
	})
}
