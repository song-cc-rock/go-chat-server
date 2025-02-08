package api

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID   string `gorm:"type:varchar(50);not null;primary_key"` // ID
	Name string `gorm:"type:varchar(255);not null;"`           // NAME
}

// TableName set table name
func (u *User) TableName() string {
	return "user"
}
