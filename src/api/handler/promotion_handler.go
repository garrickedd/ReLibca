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

// Postdata godoc
// @Summary      Create a new promotion
// @Description  API để tạo khuyến mãi mới
// @Tags         promotions
// @Accept       json
// @Produce      json
// @Param        promotion  body      model.Promotion  true  "Promotion object"
// @Success      200        {object}  map[string]interface{} "Promotion created successfully"
// @Failure      400        {object}  map[string]string "Bad Request"
// @Router       /promotion [post]
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

// Updatedata godoc
// @Summary      Update a promotion
// @Description  API để cập nhật khuyến mãi
// @Tags         promotions
// @Accept       json
// @Produce      json
// @Param        promotion  body      model.Promotion  true  "Promotion object"
// @Success      200        {object}  map[string]interface{} "Promotion updated successfully"
// @Failure      400        {object}  map[string]string "Bad Request"
// @Router       /promotion [put]
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

// Deletedata godoc
// @Summary      Delete a promotion
// @Description  API để xóa khuyến mãi
// @Tags         promotions
// @Accept       json
// @Produce      json
// @Param        promotion  body      model.Promotion  true  "Promotion object"
// @Success      200        {object}  map[string]interface{} "Promotion deleted successfully"
// @Failure      400        {object}  map[string]string "Bad Request"
// @Router       /promotion [delete]
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

// Getdata godoc
// @Summary      Get promotions
// @Description  API để lấy danh sách khuyến mãi với các tùy chọn tìm kiếm và phân trang
// @Tags         promotions
// @Accept       json
// @Produce      json
// @Param        search   query     string  false  "Search keyword"
// @Param        orderby  query     string  false  "Order by field"
// @Param        page     query     int     false  "Page number"
// @Param        limit    query     int     false  "Limit results per page"
// @Success      200      {object}  map[string]interface{} "List of promotions with pagination"
// @Failure      400      {object}  map[string]string "Bad Request"
// @Router       /promotion [get]
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
