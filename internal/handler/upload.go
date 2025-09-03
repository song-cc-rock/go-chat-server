package handler

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	v1 "go-chat-server/api/v1"
	"go-chat-server/internal/model"
	"go-chat-server/internal/service"
	"go-chat-server/pkg/s3"
	"net/http"
)

type UploadHandler struct {
	uploadService service.UploadService
	chatService   service.ChatService
}

func NewUploadHandler() *UploadHandler {
	return &UploadHandler{
		uploadService: service.NewUploadService(),
		chatService:   service.NewChatService(),
	}
}

// Upload 上传文件
func (c *UploadHandler) Upload(ctx *gin.Context) {
	// 获取参数
	file, err := ctx.FormFile("file")
	tmpId := ctx.PostForm("tmpId")
	fmt.Printf(tmpId)
	if err != nil {
		v1.HandleError(ctx, http.StatusInternalServerError, "file not found: "+err.Error())
		return
	}

	// 打开文件流
	f, err := file.Open()
	if err != nil {
		v1.HandleError(ctx, http.StatusInternalServerError, "cannot open file: "+err.Error())
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

	// 文件信息入库
	fileInfo := &model.File{}
	fileInfo.Name = objectName
	fileInfo.Size = info.Size
	fileInfo.Type = s3.GetFileExt(objectName)
	fileInfo.Path = info.Key
	fileId, dbErr := c.uploadService.SaveFileInfo(fileInfo)
	if dbErr != nil {
		v1.HandleError(ctx, http.StatusInternalServerError, "save file info error: "+dbErr.Error())
		return
	}
	// 消息文件临时ID => 实际文件ID
	dbErr1 := c.chatService.UpdateMsgFileId(tmpId, fileId)
	if dbErr1 != nil {
		v1.HandleError(ctx, http.StatusInternalServerError, "modify msg file id error: "+dbErr1.Error())
		return
	}

	v1.HandleSuccess(ctx, gin.H{
		"id": fileId,
	})
}
