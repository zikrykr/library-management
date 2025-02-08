package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zikrykr/library-management/services/book/internal/books/payload"
	"github.com/zikrykr/library-management/services/book/internal/books/port"
	"github.com/zikrykr/library-management/shared/pkg"
)

type BookHandler struct {
	bookService port.IBookService
}

func NewBookHandler(service port.IBookService) port.IBookHandler {
	return BookHandler{
		bookService: service,
	}
}

func (h BookHandler) CreateBook(c *gin.Context) {
	var data payload.CreateBookReq
	if err := c.ShouldBindJSON(&data); err != nil {
		pkg.ResponseError(c, http.StatusBadRequest, err)
		return
	}

	ctx := c.Request.Context()

	err := h.bookService.CreateBook(ctx, data)
	if err != nil {
		pkg.ResponseError(c, http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, pkg.HTTPResponse{
		Success: true,
		Message: "Successfully create book",
	})
}

func (h BookHandler) UpdateBook(c *gin.Context) {
	var data payload.UpdateBookReq
	if err := c.ShouldBindJSON(&data); err != nil {
		pkg.ResponseError(c, http.StatusBadRequest, err)
		return
	}

	ctx := c.Request.Context()

	id := c.Param("id")

	err := h.bookService.UpdateBook(ctx, id, data)
	if err != nil {
		pkg.ResponseError(c, http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, pkg.HTTPResponse{
		Success: true,
		Message: "Successfully update book",
	})
}

func (h BookHandler) GetBooks(c *gin.Context) {
	var req payload.GetBooksReq
	if err := c.ShouldBindQuery(&req); err != nil {
		pkg.ResponseError(c, http.StatusBadRequest, err)
		return
	}

	ctx := c.Request.Context()

	books, pagination, err := h.bookService.GetBooks(ctx, req)
	if err != nil {
		pkg.ResponseError(c, http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, pkg.HTTPResponse{
		Success:    true,
		Message:    "Successfully get books",
		Data:       books,
		Pagination: pagination,
	})
}

func (h BookHandler) GetBookByID(c *gin.Context) {
	ctx := c.Request.Context()

	id := c.Param("id")

	book, err := h.bookService.GetBookByID(ctx, id)
	if err != nil {
		pkg.ResponseError(c, http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, pkg.HTTPResponse{
		Success: true,
		Message: "Successfully get book",
		Data:    book,
	})
}

func (h BookHandler) DeleteBookByID(c *gin.Context) {
	ctx := c.Request.Context()

	id := c.Param("id")

	err := h.bookService.DeleteBookByID(ctx, id)
	if err != nil {
		pkg.ResponseError(c, http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, pkg.HTTPResponse{
		Success: true,
		Message: "Successfully delete book",
	})
}
