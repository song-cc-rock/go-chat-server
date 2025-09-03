package db

import (
	"go-chat-server/internal/model"
	"go-chat-server/pkg/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var DB *gorm.DB

// InitDB init db
func InitDB() {
	var err error
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n[GORM SQL] ", log.LstdFlags), // 自定义日志格式
		logger.Config{
			SlowThreshold:             time.Second, // 慢 SQL 阈值
			LogLevel:                  logger.Warn, // 日志级别：Silent, Error, Warn, Info
			IgnoreRecordNotFoundError: true,        // 忽略 ErrRecordNotFound 错误
			Colorful:                  true,        // 是否带颜色
		},
	)

	DB, err = gorm.Open(mysql.Open(config.GetString("db.mysql.dsn")), &gorm.Config{
		Logger: newLogger,
	})

	if err != nil {
		panic("❌ Failed to connect chat db")
	}

	db, _ := DB.DB()
	db.SetMaxOpenConns(config.GetInt("db.mysql.maxOpenConn"))
	db.SetMaxIdleConns(config.GetInt("db.mysql.maxIdleConn"))
	db.SetConnMaxLifetime(30 * time.Minute)
	db.SetConnMaxIdleTime(30 * time.Minute)

	log.Println("✅ Connect to chat db success")

	// init table
	InitTable(DB)
}

func InitTable(db *gorm.DB) {
	log.Println("✅ Auto migrate tables")
	err := db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&model.User{}, &model.Message{}, &model.Conversation{}, &model.File{})
	if err != nil {
		return
	}
}
