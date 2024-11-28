package api

import (
	"fmt"

	"github.com/garickedd/ReLibca/src/api/config"
	"github.com/garickedd/ReLibca/src/api/router"
	"github.com/garickedd/ReLibca/src/application/validation"
	"github.com/garickedd/ReLibca/src/shared/logging"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

var logger = logging.NewLogger(config.GetConfig())

func InitServer(cfg *config.Config) {
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())
	RegisterValidators()

	RegisterRoutes(r, cfg)
	logger := logging.NewLogger(cfg)
	logger.Info(logging.General, logging.Startup, "Started", nil)
	err := r.Run(fmt.Sprintf(":%s", cfg.Server.InternalPort))
	if err != nil {
		logger.Fatal(logging.General, logging.Startup, err.Error(), nil)
	}
}

func RegisterRoutes(r *gin.Engine, cfg *config.Config) {
	api := r.Group("/api")

	v1 := api.Group("/v1")
	{
		// User
		users := v1.Group("/users")

		// User
		router.User(users, cfg)
	}
}

func RegisterValidators() {
	val, ok := binding.Validator.Engine().(*validator.Validate)
	if ok {
		err := val.RegisterValidation("mobile", validation.VietnameseMobileNumberValidator, true)
		if err != nil {
			logger.Error(logging.Validation, logging.Startup, err.Error(), nil)
		}
		err = val.RegisterValidation("password", validation.PasswordValidator, true)
		if err != nil {
			logger.Error(logging.Validation, logging.Startup, err.Error(), nil)
		}
	}
}
