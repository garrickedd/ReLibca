package handler

import (
	"net/http"
	"strconv"

	"github.com/garrickedd/ReLibca/src/domain/model"
	"github.com/garrickedd/ReLibca/src/domain/repository"
	"github.com/gin-gonic/gin"
)

type HandlerPromotion struct {
	repository.RepoPromotionIF
}

func NewPromotion(r repository.RepoPromotionIF) *HandlerPromotion {
	return &HandlerPromotion{r}
}

func (h *HandlerPromotion) Postdata(ctx *gin.Context) {
	var promotion model.Promotion
	if err := ctx.ShouldBind(&promotion); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	response, er := h.CreatePromotion(&promotion)
	if er != nil {
		ctx.AbortWithError(http.StatusBadRequest, er)
		return
	}

	ctx.JSON(200, response)
}

func (h *HandlerPromotion) Updatedata(ctx *gin.Context) {
	var promotion model.Promotion
	if err := ctx.ShouldBind(&promotion); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	response, er := h.UpdatePromotion(&promotion)
	if er != nil {
		ctx.AbortWithError(http.StatusBadRequest, er)
		return
	}

	ctx.JSON(200, response)
}

func (h *HandlerPromotion) Deletedata(ctx *gin.Context) {
	var promotion model.Promotion
	if err := ctx.ShouldBind(&promotion); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	response, er := h.DeletePromotion(&promotion)
	if er != nil {
		ctx.AbortWithError(http.StatusBadRequest, er)
		return
	}

	ctx.JSON(200, response)
}

func (h *HandlerPromotion) Getdata(ctx *gin.Context) {
	var promotion model.Promotion
	search := ctx.Query("search")
	orderby := ctx.Query("orderby")
	page, _ := strconv.Atoi(ctx.Query("page"))
	limit, _ := strconv.Atoi(ctx.Query("limit"))

	if err := ctx.ShouldBind(&promotion); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":      http.StatusBadRequest,
			"description": "Bad Request",
			"message":     err.Error(),
		})
		return
	}

	response, pgnt, err := h.Get_Promotions(&promotion, search, limit, page, orderby)

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
		"meta":        pgnt,
	})
}
