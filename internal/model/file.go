package model

type File struct {
	ID   string `gorm:"type:varchar(50);not null;primarykey" json:"id"` // ID
	Name string `gorm:"type:varchar(64);not null;" json:"name"`         // 名称
	Size int64  `gorm:"type:bigint;not null;" json:"size"`              // 大小
	Type string `gorm:"type:varchar(64);not null;" json:"type"`         // 类型
	Path string `gorm:"type:varchar(255);not null;" json:"path"`        // 存储路径
}

func (f *File) TableName() string {
	return "file"
}
