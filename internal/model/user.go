package model

type User struct {
	ID       string `gorm:"type:varchar(50);not null;primarykey" json:"id"` // ID
	Name     string `gorm:"type:varchar(255);not null;" json:"name"`        // 账户名
	NickName string `gorm:"type:varchar(255);not null;" json:"nickName"`    // 用户名
	Phone    string `gorm:"type:varchar(50);" json:"phone"`                 // 电话
	Mail     string `gorm:"type:varchar(50);" json:"mail"`                  // 邮箱
	Password string `gorm:"type:varchar(255);not null" json:"password"`     // 密码
	Avatar   string `gorm:"type:varchar(1000);" json:"avatar"`              // 头像
	GithubId int64  `gorm:"type:int(11);" json:"githubId"`                  // Github ID
}

// TableName set table name
func (u *User) TableName() string {
	return "user"
}
