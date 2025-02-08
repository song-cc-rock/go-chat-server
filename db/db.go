package db

import (
	"github.com/spf13/viper"
	"go-chat-server/api"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

var DB *gorm.DB

// InitDB init db
func InitDB(appViper *viper.Viper) {
	var err error
	DB, err = gorm.Open(mysql.Open(appViper.GetString("db.mysql.dsn")), &gorm.Config{})

	if err != nil {
		panic("Failed to connect chat db")
	}

	db, _ := DB.DB()
	db.SetMaxOpenConns(appViper.GetInt("db.mysql.maxOpenConn"))
	db.SetMaxIdleConns(appViper.GetInt("db.mysql.maxIdleConn"))
	db.SetConnMaxLifetime(30 * time.Minute)
	db.SetConnMaxIdleTime(30 * time.Minute)

	log.Println("Connect to chat db success")

	log.Println("Auto migrate tables")
	DB.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&api.User{})
}
