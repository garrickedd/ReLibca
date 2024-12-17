package router

import (
	"github.com/garrickedd/ReLibca/src/api/middleware"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func NewRoute(db *sqlx.DB) *gin.Engine {
	router := gin.Default()

	router.Use(middleware.CORS())
	User(router, db)
	Product(router, db)
	Book(router, db)
	Promotion(router, db)
	Order(router, db)

	Auth(router, db)

	return router
}
