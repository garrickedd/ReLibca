package middleware

import (
	"strings"

	"github.com/garrickedd/ReLibca/src/infrastructure/service"
	"github.com/garrickedd/ReLibca/src/shared/config"
	"github.com/gin-gonic/gin"
)

func AuthJwt(role ...int) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var valid bool
		var header string

		if header = ctx.GetHeader("Authorization"); header == "" {
			service.NewRes(401, &config.Result{
				Message: "Please login",
			}).Send(ctx)
			return
		}

		if !strings.Contains(header, "Bearer") {
			service.NewRes(401, &config.Result{
				Message: "Header Invalid",
			}).Send(ctx)
			return
		}

		tokens := strings.Replace(header, "Bearer ", "", -1)

		check, err := service.Verifytoken(tokens)
		if err != nil {
			service.NewRes(401, &config.Result{
				Message: err.Error(),
			}).Send(ctx)
			return
		}

		for _, r := range role {
			if r == check.Role {
				valid = true
			}
		}

		if !valid {
			service.NewRes(401, &config.Result{
				Message: "Not authorized",
			}).Send(ctx)
			return
		}
		ctx.Set("userID", check.Id)
		ctx.Next()
	}
}
