package db

import (
	"go-chat-server/config"
	"go-chat-server/internal/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

var DB *gorm.DB

// InitDB init db
func InitDB() {
	var err error
	DB, err = gorm.Open(mysql.Open(config.GetString("db.mysql.dsn")), &gorm.Config{})

	if err != nil {
		panic("Failed to connect chat db")
	}

	db, _ := DB.DB()
	db.SetMaxOpenConns(config.GetInt("db.mysql.maxOpenConn"))
	db.SetMaxIdleConns(config.GetInt("db.mysql.maxIdleConn"))
	db.SetConnMaxLifetime(30 * time.Minute)
	db.SetConnMaxIdleTime(30 * time.Minute)

	log.Println("Connect to chat db success")
	log.Println("Auto migrate tables")

	// init table
	InitTable(DB)
}

func InitTable(db *gorm.DB) {
	log.Println("Migrate user table")
	err := db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&model.User{}, &model.Group{})
	if err != nil {
		return
	}
}
