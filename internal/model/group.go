package model

type Group struct {
	ID   string `gorm:"type:varchar(50);not null;primarykey"` // id
	Name string `gorm:"type:varchar(255);not null;"`          // name
}

// TableName set table name
func (u *Group) TableName() string {
	return "group"
}
