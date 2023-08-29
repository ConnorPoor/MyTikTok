package main

import (
	"github.com/ConnorPoor/MyTikTok/database"
	"github.com/ConnorPoor/MyTikTok/router"
	"github.com/gin-gonic/gin"
)

func main() {
	database.InitMysql()
	r := gin.Default()
	router.InitRouter(r)
	r.Run()
}
