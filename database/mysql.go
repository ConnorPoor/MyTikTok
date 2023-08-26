package database

import (
	"fmt"
	"github.com/ConnorPoor/MyTikTok/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	userName string = "normal"
	passWord string = "123"
	url      string = "127.0.0.1"
	port     string = "3306"
	dbName   string = "MyTikTok"

	db *gorm.DB
)

func InitMysql() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		userName,
		passWord,
		url,
		port,
		dbName,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Favorite{})
	db.AutoMigrate(&model.Following{})
	db.AutoMigrate(&model.Followers{})
	db.AutoMigrate(&model.Comment{})
	db.AutoMigrate(&model.Video{})
}
