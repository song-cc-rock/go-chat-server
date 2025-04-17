package model

type Message struct {
	ID        string `gorm:"type:varchar(50);not null;primarykey"` // id
	FromId    string `gorm:"type:varchar(64);not null;"`
	ToId      string `gorm:"type:varchar(64);not null;"`
	Content   string `gorm:"type:text;not null;"`
	MsgType   string `gorm:"type:varchar(20);default:'text';"` // text, image, video
	Status    string `gorm:"type:varchar(20);default:'sent';"` // sent, read, deleted
	CreatedAt int64  `gorm:"type:bigint;not null;"`            // 创建时间
}
