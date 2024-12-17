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

// AddOrder godoc
// @Summary      Create a new order
// @Description  API để thêm đơn hàng mới
// @Tags         orders
// @Accept       json
// @Produce      json
// @Param        order  body      model.Order  true  "Order object"
// @Success      200    {object}  map[string]interface{} "Order created successfully"
// @Failure      400    {object}  map[string]string "Bad Request"
// @Router       /order [post]
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

// RemoveOrder godoc
// @Summary      Delete an order
// @Description  API để xóa đơn hàng
// @Tags         orders
// @Accept       json
// @Produce      json
// @Param        order  body      model.Order  true  "Order object"
// @Success      200    {object}  map[string]interface{} "Order deleted successfully"
// @Failure      400    {object}  map[string]string "Bad Request"
// @Router       /order [delete]
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
