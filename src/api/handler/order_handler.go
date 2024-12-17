package handler

import (
	"net/http"

	"github.com/garrickedd/ReLibca/src/domain/model"
	"github.com/garrickedd/ReLibca/src/domain/repository"
	"github.com/gin-gonic/gin"
)

type HandlerOrder struct {
	repository.RepoOrderIF
}

func NewOrder(r repository.RepoOrderIF) *HandlerOrder {
	return &HandlerOrder{r}
}

func (h *HandlerOrder) AddOrder(ctx *gin.Context) {
	var order model.Order
	if err := ctx.ShouldBind(&order); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	response, er := h.CreateOrder(&order)
	if er != nil {
		ctx.AbortWithError(http.StatusBadRequest, er)
		return
	}

	ctx.JSON(200, response)
}

func (h *HandlerOrder) RemoveOrder(ctx *gin.Context) {
	var order model.Order
	if err := ctx.ShouldBind(&order); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	response, er := h.DeleteOrder(&order)
	if er != nil {
		ctx.AbortWithError(http.StatusBadRequest, er)
		return
	}

	ctx.JSON(200, response)
}
