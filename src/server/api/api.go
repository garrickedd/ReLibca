package api

import (
	"fmt"

	"github.com/garrickedd/ReLibca/src/server/api/routers"
	"github.com/garrickedd/ReLibca/src/server/config"
	"github.com/gin-gonic/gin"
)

func InitServer() {
	cfg := config.GetConfig()
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())

	v1 := r.Group("/api/v1/")
	{
		health := v1.Group("/health")
		routers.Health(health)
	}

	if err := r.Run(fmt.Sprintf(":%s", cfg.Server.InternalPort)); err != nil {
		panic(err)
	} // Added
}
