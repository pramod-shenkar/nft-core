package db

import (
	"nftledger/internal/core/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Conn struct {
	*gorm.DB
}

func New() (*Conn, error) {
	dsn := "user:password@tcp(127.0.0.1:3306)/nftledger?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&model.Request{})

	return &Conn{
		db,
	}, nil
}
