package model

type Friend struct {
	ID        string `gorm:"type:varchar(50);not null;primarykey"` // ID
	UserId    string `gorm:"type:varchar(50);not null;index"`      // 用户ID
	FriendId  string `gorm:"type:varchar(50);not null;index"`      // 好友ID
	Status    string `gorm:"type:varchar(20);default:'pending';"`  // 好友状态: pending => 待确认, accepted => 已添加, blocked => 已拉黑
	CreatedAt int64  `gorm:"type:bigint;not null;"`                // 创建时间
}

func (f *Friend) TableName() string {
	return "friend"
}
