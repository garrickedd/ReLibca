package handler

import (
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

// Postdata godoc
// @Summary      Create a new book
// @Description  API để tạo sách mới
// @Tags         books
// @Accept       json
// @Produce      json
// @Param        book  body      model.Book  true  "Book object"
// @Success      200   {object}  map[string]interface{} "Book created successfully"
// @Failure      400   {object}  map[string]string "Bad Request"
// @Router       /book [post]
func (h *HandlerBook) Postdata(ctx *gin.Context) {
	var book model.Book
	if err := ctx.ShouldBind(&book); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// book.Image_file = ctx.MustGet("image").(string)
	// fmt.Println(book)
	response, er := h.CreateBook(&book)
	if er != nil {
		ctx.AbortWithError(http.StatusBadRequest, er)
		return
	}

	ctx.JSON(200, response)
}

// Updatedata godoc
// @Summary      Update a book
// @Description  API để cập nhật thông tin sách
// @Tags         books
// @Accept       json
// @Produce      json
// @Param        book  body      model.Book  true  "Book object"
// @Success      200   {object}  map[string]interface{} "Book updated successfully"
// @Failure      400   {object}  map[string]string "Bad Request"
// @Router       /book [put]
func (h *HandlerBook) Updatedata(ctx *gin.Context) {
	bookTitle := ctx.Param("title") //

	var book model.Book
	if err := ctx.ShouldBind(&book); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	book.Title = bookTitle

	response, er := h.UpdateBook(&book)
	if er != nil {
		ctx.AbortWithError(http.StatusBadRequest, er)
		return
	}

	ctx.JSON(200, response)
}

// Deletedata godoc
// @Summary      Delete a book
// @Description  API để xóa sách
// @Tags         books
// @Accept       json
// @Produce      json
// @Param        book  body      model.Book  true  "Book object"
// @Success      200   {object}  map[string]interface{} "Book deleted successfully"
// @Failure      400   {object}  map[string]string "Bad Request"
// @Router       /book [delete]
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

// Getdata godoc
// @Summary      Get books
// @Description  API để lấy danh sách sách với các tùy chọn tìm kiếm và phân trang
// @Tags         books
// @Accept       json
// @Produce      json
// @Param        search   query     string  false  "Search keyword"
// @Param        orderby  query     string  false  "Order by field"
// @Param        page     query     int     false  "Page number"
// @Param        limit    query     int     false  "Limit results per page"
// @Success      200      {object}  map[string]interface{} "List of books with pagination"
// @Failure      400      {object}  map[string]string "Bad Request"
// @Router       /book [get]
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
