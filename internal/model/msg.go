package model

type Message struct {
	ID        string `gorm:"type:varchar(50);not null;primarykey"` // ID
	FromId    string `gorm:"type:varchar(64);not null;"`           // 发送者ID
	ToId      string `gorm:"type:varchar(64);not null;"`           // 接收者ID
	Content   string `gorm:"type:text;not null;"`                  // 消息内容
	MsgType   string `gorm:"type:varchar(20);default:'text';"`     // 消息类型: text, image, video
	Status    string `gorm:"type:varchar(20);default:'sent';"`     // 消息状态: sent, read, deleted
	CreatedAt int64  `gorm:"type:bigint;not null;"`                // 创建时间
}

func (u *Message) TableName() string {
	return "messages"
}
