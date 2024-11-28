package api

import (
	"github.com/garickedd/ReLibca/src/api/config"
	"github.com/garickedd/ReLibca/src/shared/logging"
	"github.com/gin-gonic/gin"
)

func InitServer(cfg *config.Config) {
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())

	// logging
	logger := logging.NewLogger(cfg)
	logger.Info(logging.General, logging.Startup, "Started", nil)
	r.Run(":5005")
}

// func RegisterRoutes(r *gin.Engine, cfg *config.Config) {
// 	api := r.Group("/api")

// 	v1 := api.Group("/v1")
// 	{
// 		// Product

// 	}
// }
