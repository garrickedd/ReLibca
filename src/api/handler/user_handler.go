package handler

import (
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/garrickedd/ReLibca/src/domain/model"
	"github.com/garrickedd/ReLibca/src/domain/repository"
	"github.com/garrickedd/ReLibca/src/infrastructure/service"
	"github.com/garrickedd/ReLibca/src/shared/config"
	"github.com/gin-gonic/gin"
)

type HandlerUser struct {
	repository.RepoUserIF
}

func NewUser(r repository.RepoUserIF) *HandlerUser {
	return &HandlerUser{r}
}

func (h *HandlerUser) Postdata(ctx *gin.Context) {
	var err error
	user := model.User{
		Role: 0,
	}
	if err := ctx.ShouldBind(&user); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	_, err = govalidator.ValidateStruct(&user)
	if err != nil {
		service.NewRes(401, &config.Result{
			Data: err.Error(),
		}).Send(ctx)
		return
	}

	// user.Image_file = ctx.GetString("image")
	// user.Pass, err = service.HashPassword(user.Pass)
	// if err != nil {
	// 	pkg.NewRes(401, &config.Result{
	// 		Data: err.Error(),
	// 	}).Send(ctx)
	// 	return
	// }

	response, er := h.CreateUser(&user)
	if er != nil {
		ctx.AbortWithError(http.StatusBadRequest, er)
		return
	}

	ctx.JSON(200, response)
}

func (h *HandlerUser) Updatedata(ctx *gin.Context) {
	var user model.User
	user.Username = ctx.Param("username")
	if err := ctx.ShouldBind(&user); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	user.Phone = ctx.PostForm("phone")
	response, er := h.UpdateUser(&user)
	if er != nil {
		ctx.AbortWithError(http.StatusBadRequest, er)
		return
	}

	ctx.JSON(200, response)
}

func (h *HandlerUser) Deletedata(ctx *gin.Context) {
	var user model.User
	if err := ctx.ShouldBind(&user); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	response, er := h.DeleteUser(&user)
	if er != nil {
		ctx.AbortWithError(http.StatusBadRequest, er)
		return
	}

	ctx.JSON(200, response)
}

func (h *HandlerUser) Getdata(ctx *gin.Context) {
	var users model.User
	search := ctx.Query("search")
	if err := ctx.ShouldBind(&users); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":      http.StatusBadRequest,
			"description": "Bad Request",
			"message":     err.Error(),
		})
		return
	}

	response, err := h.Get_Users(&users, search)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":      http.StatusBadRequest,
			"description": "Bad Request",
			"message":     err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":      http.StatusOK,
		"description": "OK",
		"data":        response,
	})
}
