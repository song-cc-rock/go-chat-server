package model

type UserRefreshToken struct {
	ID           string `gorm:"type:varchar(50);not null;primarykey"`   // ID
	UserID       string `gorm:"type:varchar(50);not null;index"`        // 所属用户
	RefreshToken string `gorm:"type:varchar(255);not null;uniqueIndex"` // refresh token，本质是随机字符串或JWT
	ExpiresAt    int64  `gorm:"type:bigint;not null"`                   // 过期时间
	Revoked      bool   `gorm:"default:false"`                          // 是否已失效（用户登出/管理员强制下线）
	CreatedAt    int64  `gorm:"type:bigint;not null;"`                  // 创建时间
	UpdatedAt    int64  `gorm:"type:bigint;not null;"`                  // 更新时间
}

func (u *UserRefreshToken) TableName() string {
	return "refresh_token"
}
