package handler

import (
	"github.com/garrickedd/ReLibca/src/domain/repository"
	"github.com/garrickedd/ReLibca/src/infrastructure/service"
	"github.com/garrickedd/ReLibca/src/shared/config"
	"github.com/gin-gonic/gin"
)

type User struct {
	Username string `db:"username" json:"username" form:"username"`
	Pass     string `db:"pass" json:"pass,omitempty"`
}

type HandlerAuth struct {
	*repository.RepoUser
}

func NewAuth(r *repository.RepoUser) *HandlerAuth {
	return &HandlerAuth{r}
}

func (h *HandlerAuth) Login(ctx *gin.Context) {
	var data *User
	if ers := ctx.ShouldBind(&data); ers != nil {
		service.NewRes(500, &config.Result{
			Data: ers.Error(),
		}).Send(ctx)
		return
	}

	users, err := h.GetAuthUser(data.Username)
	if err != nil {
		service.NewRes(401, &config.Result{
			Data: err.Error(),
		}).Send(ctx)
		return
	}

	if err := service.VerifyPassword(users.Pass, data.Pass); err != nil {
		service.NewRes(401, &config.Result{
			Data: "Wrong password",
		}).Send(ctx)
		return
	}

	jsonToken := service.NewToken(users.Id_user, users.Role)
	tokens, err := jsonToken.Generate()
	if err != nil {
		service.NewRes(500, &config.Result{
			Data: err.Error(),
		}).Send(ctx)
		return
	}

	service.NewRes(200, &config.Result{Data: tokens}).Send(ctx)
}
