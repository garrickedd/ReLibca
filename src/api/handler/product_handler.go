package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/garrickedd/ReLibca/src/domain/model"
	"github.com/garrickedd/ReLibca/src/domain/repository"
	"github.com/gin-gonic/gin"
)

type HandlerProduct struct {
	repository.RepoProductIF
}

func NewProduct(r repository.RepoProductIF) *HandlerProduct {
	return &HandlerProduct{r}
}

func (h *HandlerProduct) Postdata(ctx *gin.Context) {
	var product model.Product
	if err := ctx.ShouldBind(&product); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	product.Image_file = ctx.MustGet("image").(string)
	fmt.Println(product)
	response, er := h.CreateProduct(&product)
	if er != nil {
		ctx.AbortWithError(http.StatusBadRequest, er)
		return
	}

	ctx.JSON(200, response)
}

func (h *HandlerProduct) Updatedata(ctx *gin.Context) {
	var product model.Product
	if err := ctx.ShouldBind(&product); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	response, er := h.UpdateProduct(&product)
	if er != nil {
		ctx.AbortWithError(http.StatusBadRequest, er)
		return
	}

	ctx.JSON(200, response)
}

func (h *HandlerProduct) Deletedata(ctx *gin.Context) {
	var product model.Product
	if err := ctx.ShouldBind(&product); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	response, er := h.DeleteProduct(&product)
	if er != nil {
		ctx.AbortWithError(http.StatusBadRequest, er)
		return
	}

	ctx.JSON(200, response)
}

func (h *HandlerProduct) Getdata(ctx *gin.Context) {
	var product model.Product
	search := ctx.Query("search")
	orderby := ctx.Query("orderby")
	page, _ := strconv.Atoi(ctx.Query("page"))
	limit, _ := strconv.Atoi(ctx.Query("limit"))

	if err := ctx.ShouldBind(&product); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":      http.StatusBadRequest,
			"description": "Bad Request",
			"message":     err.Error(),
		})
		return
	}

	response, pgnt, err := h.Get_Products(&product, search, limit, page, orderby)

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
