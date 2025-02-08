package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zikrykr/library-management/services/author/internal/authors/payload"
	"github.com/zikrykr/library-management/services/author/internal/authors/port"
	"github.com/zikrykr/library-management/shared/pkg"
)

type AuthorHandler struct {
	authorService port.IAuthorService
}

func NewAuthorHandler(service port.IAuthorService) port.IAuthorHandler {
	return AuthorHandler{
		authorService: service,
	}
}

func (h AuthorHandler) CreateAuthor(c *gin.Context) {
	var data payload.CreateAuthorReq
	if err := c.ShouldBindJSON(&data); err != nil {
		pkg.ResponseError(c, http.StatusBadRequest, err)
		return
	}

	ctx := c.Request.Context()

	err := h.authorService.CreateAuthor(ctx, data)
	if err != nil {
		pkg.ResponseError(c, http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, pkg.HTTPResponse{
		Success: true,
		Message: "Successfully create author",
	})
}

func (h AuthorHandler) UpdateAuthor(c *gin.Context) {
	var data payload.UpdateAuthorReq
	if err := c.ShouldBindJSON(&data); err != nil {
		pkg.ResponseError(c, http.StatusBadRequest, err)
		return
	}

	ctx := c.Request.Context()

	id := c.Param("id")

	err := h.authorService.UpdateAuthor(ctx, id, data)
	if err != nil {
		pkg.ResponseError(c, http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, pkg.HTTPResponse{
		Success: true,
		Message: "Successfully update author",
	})
}

func (h AuthorHandler) GetAuthors(c *gin.Context) {
	var req payload.GetAuthorsReq
	if err := c.ShouldBindQuery(&req); err != nil {
		pkg.ResponseError(c, http.StatusBadRequest, err)
		return
	}

	ctx := c.Request.Context()

	authors, pagination, err := h.authorService.GetAuthors(ctx, req)
	if err != nil {
		pkg.ResponseError(c, http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, pkg.HTTPResponse{
		Success:    true,
		Message:    "Successfully get authors",
		Data:       authors,
		Pagination: pagination,
	})
}

func (h AuthorHandler) GetAuthorByID(c *gin.Context) {
	ctx := c.Request.Context()

	id := c.Param("id")

	author, err := h.authorService.GetAuthorByID(ctx, id)
	if err != nil {
		pkg.ResponseError(c, http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, pkg.HTTPResponse{
		Success: true,
		Message: "Successfully get author",
		Data:    author,
	})
}

func (h AuthorHandler) DeleteAuthorByID(c *gin.Context) {
	ctx := c.Request.Context()

	id := c.Param("id")

	err := h.authorService.DeleteAuthorByID(ctx, id)
	if err != nil {
		pkg.ResponseError(c, http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, pkg.HTTPResponse{
		Success: true,
		Message: "Successfully delete author",
	})
}
