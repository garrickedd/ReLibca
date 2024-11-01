package routes

import (
	"github.com/garrickedd/ReLibca/src/api/controllers"
	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {
	router.GET("/", controllers.GetUser)
	// router.POST("/", controllers.CreateUser)
	// router.DELETE("/:id", controllers.DeleteUser)
	// router.PUT("/:id", controllers.UpdateUser)
}