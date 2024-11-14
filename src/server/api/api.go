package api

import (
	"fmt"

	"github.com/garrickedd/ReLibca/src/server/api/routers"
	validation "github.com/garrickedd/ReLibca/src/server/api/validations"
	"github.com/garrickedd/ReLibca/src/server/config"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func InitServer(cfg *config.Config) {
	r := gin.New()
	// r.Use(gin.Logger(), gin.Recovery())

	RegisterValidators()
	// val, ok := binding.Validator.Engine().(*validator.Validate)
	// if ok {
	// 	val.RegisterValidation("mobile", validation.VietnameseMobileNumberValidator, true)
	// }

	// TODO: using cors middleware
	r.Use(gin.Logger(), gin.Recovery())

	api := r.Group("/api")

	v1 := api.Group("/v1/")
	{
		health := v1.Group("/health")
		test_router := v1.Group("/test")

		routers.Health(health)
		routers.TestRouter(test_router)
	}

	v2 := api.Group("/v2")
	{
		health := v2.Group("/health")
		routers.Health(health)
	}

	if err := r.Run(fmt.Sprintf(":%s", cfg.Server.InternalPort)); err != nil {
		panic(err)
	} // Added
}

func RegisterValidators() {
	val, ok := binding.Validator.Engine().(*validator.Validate)
	if ok {
		val.RegisterValidation("mobile", validation.VietnameseMobileNumberValidator, true)
		// val.RegisterValidation("password", validation.PasswordValidator, true)
	}
}
