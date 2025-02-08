package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/zikrykr/library-management/services/category/internal/categories/port"
)

type (
	categoryRoutes struct{}
)

var (
	Routes categoryRoutes
)

func (r categoryRoutes) NewRoutes(router *gin.RouterGroup, categoryHandler port.ICategoryHandler) {
	// get categories
	router.GET("", categoryHandler.GetCategories)
	// get categories by id
	router.GET("/:id", categoryHandler.GetCategoryByID)
	// create categories
	router.POST("", categoryHandler.CreateCategory)
	// update categories
	router.PUT("/:id", categoryHandler.UpdateCategory)
	// delete categories
	router.DELETE("/:id", categoryHandler.DeleteCategoryByID)
}
