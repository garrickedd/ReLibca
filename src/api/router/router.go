package router

import (
	"github.com/garrickedd/ReLibca/src/api/middleware"
	"github.com/garrickedd/ReLibca/src/docs"
	"github.com/gin-gonic/gin"

	"github.com/jmoiron/sqlx"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRoute(db *sqlx.DB) *gin.Engine {
	router := gin.Default()
	docs.SwaggerInfo.BasePath = "/"

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Use(middleware.CORS())
	User(router, db)
	Product(router, db)
	Book(router, db)
	Promotion(router, db)
	Order(router, db)

	Auth(router, db)

	return router
}
