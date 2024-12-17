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

// Postdata godoc
// @Summary      Create a new product
// @Description  API để tạo sản phẩm mới
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        product  body      model.Product  true  "Product object"
// @Success      200      {object}  map[string]interface{} "Product created successfully"
// @Failure      400      {object}  map[string]string "Bad Request"
// @Router       /product [post]
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

// Updatedata godoc
// @Summary      Update product
// @Description  API để cập nhật sản phẩm
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        product  body      model.Product  true  "Product object"
// @Success      200      {object}  map[string]interface{} "Product updated successfully"
// @Failure      400      {object}  map[string]string "Bad Request"
// @Router       /product [put]
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

// Deletedata godoc
// @Summary      Delete product
// @Description  API để xóa sản phẩm
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        product  body      model.Product  true  "Product object"
// @Success      200      {object}  map[string]interface{} "Product deleted successfully"
// @Failure      400      {object}  map[string]string "Bad Request"
// @Router       /product [delete]
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

// Getdata godoc
// @Summary      Get products
// @Description  API để lấy danh sách sản phẩm có phân trang
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        search   query     string  false  "Search keyword"
// @Param        orderby  query     string  false  "Order by field"
// @Param        page     query     int     false  "Page number"
// @Param        limit    query     int     false  "Limit results per page"
// @Success      200      {object}  map[string]interface{} "List of products with pagination"
// @Failure      400      {object}  map[string]string "Bad Request"
// @Router       /product [get]
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
