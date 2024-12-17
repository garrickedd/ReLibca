package router

import (
	"github.com/garrickedd/ReLibca/src/api/handler"
	"github.com/garrickedd/ReLibca/src/api/middleware"
	"github.com/garrickedd/ReLibca/src/domain/repository"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func Promotion(g *gin.Engine, d *sqlx.DB) {
	route := g.Group("/promotion")

	repo := repository.NewPromotion(d)
	handler := handler.NewPromotion(repo)

	route.POST("/", handler.Postdata)
	route.GET("/", handler.Getdata)
	route.PUT("/:code", middleware.AuthJwt(1), handler.Updatedata)
	route.DELETE("/", middleware.AuthJwt(1), handler.Deletedata)
}
