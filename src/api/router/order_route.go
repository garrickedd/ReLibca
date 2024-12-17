package router

import (
	"github.com/garrickedd/ReLibca/src/api/handler"
	"github.com/garrickedd/ReLibca/src/domain/repository"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func Order(g *gin.Engine, d *sqlx.DB) {
	route := g.Group("/order")

	repo := repository.NewOrder(d)
	handler := handler.NewOrder(repo)

	route.POST("/", handler.AddOrder)
	route.DELETE("/", handler.RemoveOrder)

}
