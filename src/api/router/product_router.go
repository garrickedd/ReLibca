package router

import (
	"github.com/garrickedd/ReLibca/src/api/handler"
	"github.com/garrickedd/ReLibca/src/api/middleware"
	"github.com/garrickedd/ReLibca/src/domain/repository"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func Product(g *gin.Engine, d *sqlx.DB) {
	route := g.Group("/product")

	repo := repository.NewProduct(d)
	handler := handler.NewProduct(repo)

	// orderRepo := repository.NewOrder(d)
	// orderHandler := handler.NewOrder(orderRepo)

	route.POST("/", handler.Postdata)
	// route.POST("/order", orderHandler.AddOrder)
	route.GET("/", handler.Getdata)
	route.PUT("/:product_name", middleware.AuthJwt(1), handler.Updatedata)
	route.DELETE("/", middleware.AuthJwt(1), handler.Deletedata)
}
