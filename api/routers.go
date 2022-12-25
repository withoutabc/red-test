package api

import "github.com/gin-gonic/gin"

func InitRouter() {
	r := gin.Default()
	r.POST("/register", Login)
	r.POST("/add", AddRepository)
	r.Run()
}
