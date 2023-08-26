package main

import (
	"github.com/ConnorPoor/MyTikTok/database"
	"github.com/gin-gonic/gin"
)

func main() {
	database.InitMysql()
	r := gin.Default()
	r.Run()
}
