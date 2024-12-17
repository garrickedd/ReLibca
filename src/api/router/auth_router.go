package router

import (
	"github.com/garrickedd/ReLibca/src/api/handler"
	"github.com/garrickedd/ReLibca/src/domain/repository"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func Auth(g *gin.Engine, d *sqlx.DB) {
	route := g.Group("/login")

	repo := repository.NewUser(d)
	handler := handler.NewAuth(repo)

	route.POST("/auth", handler.Login)

}
