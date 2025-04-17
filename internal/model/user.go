package model

type User struct {
	ID       string `gorm:"type:varchar(50);not null;primarykey"` // ID
	Name     string `gorm:"type:varchar(255);not null;"`          // 账户名
	NickName string `gorm:"type:varchar(255);not null;"`          // 用户名
	Phone    string `gorm:"type:int(11);"`                        // 电话
	Mail     string `gorm:"type:varchar(50);"`                    // 邮箱
	Password string `gorm:"type:varchar(255);not null"`           // 密码
	Avatar   string `gorm:"type:varchar(100);"`                   // 头像
	GithubId int64  `gorm:"type:int(11);"`                        // Github ID
}

// TableName set table name
func (u *User) TableName() string {
	return "user"
}
