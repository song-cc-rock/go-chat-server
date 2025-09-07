package model

type Friend struct {
	ID        string `gorm:"type:varchar(50);not null;primarykey"` // ID
	UserId    string `gorm:"type:varchar(50);not null;index"`      // 用户ID
	FriendId  string `gorm:"type:varchar(50);not null;index"`      // 好友ID
	Status    string `gorm:"type:varchar(20);default:'pending';"`  // 好友状态: 正常 => normal，拉黑 => blocked，删除 => deleted
	CreatedAt int64  `gorm:"type:bigint;not null;"`                // 创建时间
}

type FriendRequest struct {
	ID        string `gorm:"type:varchar(50);not null;primarykey" json:"id"`    // ID
	FromId    string `gorm:"type:varchar(50);not null;index" json:"fromId"`     // 请求发起者ID
	ToId      string `gorm:"type:varchar(50);not null;index" json:"toId"`       // 请求接收者ID
	Message   string `gorm:"type:varchar(255);" json:"message"`                 // 请求附加消息
	Status    string `gorm:"type:varchar(20);default:'pending';" json:"status"` // 请求状态: pending => 待处理, accepted => 已接受, rejected => 已拒绝
	CreatedAt int64  `gorm:"type:bigint;not null;" json:"createdAt"`            // 创建时间
}

func (f *Friend) TableName() string {
	return "friend"
}

func (fr *FriendRequest) TableName() string {
	return "friend_request"
}
