package routers

import (
	"github.com/garrickedd/ReLibca/src/server/api/handlers"
	"github.com/gin-gonic/gin"
)

func Health(r *gin.RouterGroup) {
	handler := handlers.NewHealthHandler()

	r.GET("/", handler.Health) // Call handler.Health method
}
