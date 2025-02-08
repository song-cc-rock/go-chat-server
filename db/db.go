package db

import (
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

var db *gorm.DB

// InitDB init db
func InitDB(appViper *viper.Viper) {
	var err error
	db, err = gorm.Open(mysql.Open(appViper.GetString("db.mysql.dsn")), &gorm.Config{})

	if err != nil {
		panic("Failed to connect chat db")
	}

	DB, _ := db.DB()
	DB.SetMaxOpenConns(appViper.GetInt("db.mysql.maxOpenConn"))
	DB.SetMaxIdleConns(appViper.GetInt("db.mysql.maxIdleConn"))
	DB.SetConnMaxLifetime(30 * time.Minute)
	DB.SetConnMaxIdleTime(30 * time.Minute)

	log.Println("Connect to chat db success")
}
