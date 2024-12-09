package router

import (
	"github.com/garrickedd/ReLibca/src/api/handler"
	"github.com/garrickedd/ReLibca/src/api/middleware"
	"github.com/garrickedd/ReLibca/src/domain/repository"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func User(g *gin.Engine, d *sqlx.DB) {
	route := g.Group("/user")

	repo := repository.NewUser(d)
	handler := handler.NewUser(repo)

	route.POST("/", handler.Postdata)
	route.GET("/", middleware.AuthJwt(0, 1), handler.Getdata)
	route.PUT("/:username", middleware.AuthJwt(1), handler.Updatedata)
	route.DELETE("/", middleware.AuthJwt(1), handler.Deletedata)
}
