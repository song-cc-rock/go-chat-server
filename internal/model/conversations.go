package model

type Conversation struct {
	ID            string `gorm:"type:varchar(50);not null;primarykey"` // ID
	UserID        string `gorm:"type:varchar(50);not null;"`           // 用户
	TargetUserID  string `gorm:"type:varchar(50);not null;"`           // 会话对象
	LastMessage   string `gorm:"type:text;not null;"`                  // 最后一条消息
	LastMessageAt int64  `gorm:"type:bigint;not null;"`                // 最后一条消息时间
	UnreadCount   int64  `gorm:"type:bigint;not null;"`                // 未读消息数量
	Deleted       bool   `gorm:"type:bool;default:false"`              // 是否删除
}

func (u *Conversation) TableName() string {
	return "conversation"
}
