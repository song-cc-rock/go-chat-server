package model

type User struct {
	ID       string `gorm:"type:varchar(50);not null;primarykey"` // id
	Name     string `gorm:"type:varchar(255);not null;"`          // name
	NickName string `gorm:"type:varchar(255);not null;"`          // nickname
	Phone    string `gorm:"type:int(11);"`                        // phone
	Mail     string `gorm:"type:varchar(50);"`                    // mail
	Password string `gorm:"type:varchar(255);not null"`           // password
	Avatar   string `gorm:"type:varchar(100);"`                   // avatar
}

// TableName set table name
func (u *User) TableName() string {
	return "user"
}
