package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/garrickedd/ReLibca/src/domain/model"
	"github.com/garrickedd/ReLibca/src/domain/repository"
	"github.com/gin-gonic/gin"
)

type HandlerBook struct {
	repository.RepoBookIF
}

func NewBook(r repository.RepoBookIF) *HandlerBook {
	return &HandlerBook{r}
}

func (h *HandlerBook) Postdata(ctx *gin.Context) {
	var book model.Book
	if err := ctx.ShouldBind(&book); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	book.Image_file = ctx.MustGet("image").(string)
	fmt.Println(book)
	response, er := h.CreateBook(&book)
	if er != nil {
		ctx.AbortWithError(http.StatusBadRequest, er)
		return
	}

	ctx.JSON(200, response)
}

func (h *HandlerBook) Updatedata(ctx *gin.Context) {
	var book model.Book
	if err := ctx.ShouldBind(&book); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	response, er := h.UpdateBook(&book)
	if er != nil {
		ctx.AbortWithError(http.StatusBadRequest, er)
		return
	}

	ctx.JSON(200, response)
}

func (h *HandlerBook) Deletedata(ctx *gin.Context) {
	var book model.Book
	if err := ctx.ShouldBind(&book); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	response, er := h.DeleteBook(&book)
	if er != nil {
		ctx.AbortWithError(http.StatusBadRequest, er)
		return
	}

	ctx.JSON(200, response)
}

func (h *HandlerBook) Getdata(ctx *gin.Context) {
	var book model.Book
	search := ctx.Query("search")
	orderby := ctx.Query("orderby")
	page, _ := strconv.Atoi(ctx.Query("page"))
	limit, _ := strconv.Atoi(ctx.Query("limit"))

	if err := ctx.ShouldBind(&book); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":      http.StatusBadRequest,
			"description": "Bad Request",
			"message":     err.Error(),
		})
		return
	}

	response, pgnt, err := h.Get_Books(&book, search, limit, page, orderby)

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
