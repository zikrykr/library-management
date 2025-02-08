package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zikrykr/library-management/services/category/internal/categories/payload"
	"github.com/zikrykr/library-management/services/category/internal/categories/port"
	"github.com/zikrykr/library-management/shared/pkg"
)

type CategoryHandler struct {
	categoryService port.ICategoryService
}

func NewCategoryHandler(service port.ICategoryService) port.ICategoryHandler {
	return CategoryHandler{
		categoryService: service,
	}
}

func (h CategoryHandler) CreateCategory(c *gin.Context) {
	var data payload.CreateCategoryReq
	if err := c.ShouldBindJSON(&data); err != nil {
		pkg.ResponseError(c, http.StatusBadRequest, err)
		return
	}

	ctx := c.Request.Context()

	err := h.categoryService.CreateCategory(ctx, data)
	if err != nil {
		pkg.ResponseError(c, http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, pkg.HTTPResponse{
		Success: true,
		Message: "Successfully create category",
	})
}

func (h CategoryHandler) UpdateCategory(c *gin.Context) {
	var data payload.UpdateCategoryReq
	if err := c.ShouldBindJSON(&data); err != nil {
		pkg.ResponseError(c, http.StatusBadRequest, err)
		return
	}

	ctx := c.Request.Context()

	id := c.Param("id")

	err := h.categoryService.UpdateCategory(ctx, id, data)
	if err != nil {
		pkg.ResponseError(c, http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, pkg.HTTPResponse{
		Success: true,
		Message: "Successfully update category",
	})
}

func (h CategoryHandler) GetCategories(c *gin.Context) {
	var req payload.GetCategoriesReq
	if err := c.ShouldBindQuery(&req); err != nil {
		pkg.ResponseError(c, http.StatusBadRequest, err)
		return
	}

	ctx := c.Request.Context()

	authors, pagination, err := h.categoryService.GetCategories(ctx, req)
	if err != nil {
		pkg.ResponseError(c, http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, pkg.HTTPResponse{
		Success:    true,
		Message:    "Successfully get categories",
		Data:       authors,
		Pagination: pagination,
	})
}

func (h CategoryHandler) GetCategoryByID(c *gin.Context) {
	ctx := c.Request.Context()

	id := c.Param("id")

	author, err := h.categoryService.GetCategoryByID(ctx, id)
	if err != nil {
		pkg.ResponseError(c, http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, pkg.HTTPResponse{
		Success: true,
		Message: "Successfully get category",
		Data:    author,
	})
}

func (h CategoryHandler) DeleteCategoryByID(c *gin.Context) {
	ctx := c.Request.Context()

	id := c.Param("id")

	err := h.categoryService.DeleteCategoryByID(ctx, id)
	if err != nil {
		pkg.ResponseError(c, http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, pkg.HTTPResponse{
		Success: true,
		Message: "Successfully delete category",
	})
}
