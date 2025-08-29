package bootstrap

import (
	"context"
	"go-chat-server/pkg/config"
	"go-chat-server/pkg/db"
	"go-chat-server/pkg/s3"
)

func Init() error {
	// init config
	config.Init()
	// init db
	db.InitDB()

	// init s3
	if err := s3.Init(); err != nil {
		return err
	}

	// create bucket
	if err := s3.CreateBucket(context.Background(), "go-chat"); err != nil {
		return err
	}

	return nil
}
