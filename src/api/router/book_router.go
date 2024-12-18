package router

import (
	"github.com/garrickedd/ReLibca/src/api/handler"
	"github.com/garrickedd/ReLibca/src/api/middleware"
	"github.com/garrickedd/ReLibca/src/domain/repository"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func Book(g *gin.Engine, d *sqlx.DB) {
	route := g.Group("/book")

	repo := repository.NewBook(d)
	handler := handler.NewBook(repo)

	route.POST("/", handler.Postdata)
	route.GET("/", handler.Getdata)
	route.PUT("/:title", middleware.AuthJwt(1), handler.Updatedata)
	route.DELETE("/", middleware.AuthJwt(1), handler.Deletedata)
}
