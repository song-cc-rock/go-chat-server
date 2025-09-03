package v1

type ConversationHisRequest struct {
	ConversationID string `form:"id" binding:"required"`
	Page           int    `form:"page,default=1"`
	Size           int    `form:"size,default=20"`
}
