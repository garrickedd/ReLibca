package api

import "github.com/gin-gonic/gin"

func InitServer() {
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())

	r.Run(":5005")
}
