package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/zikrykr/library-management/services/category/internal/categories/port"
)

type (
	routes      struct{}
	adminRoutes struct{}
)

var (
	Routes      routes
	AdminRoutes adminRoutes
)

func (r adminRoutes) NewAdminRoutes(router *gin.RouterGroup, categoryHandler port.ICategoryHandler) {
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

func (r routes) NewRoutes(router *gin.RouterGroup, categoryHandler port.ICategoryHandler) {
	// get categories
	router.GET("", categoryHandler.GetCategories)
	// get categories by id
	router.GET("/:id", categoryHandler.GetCategoryByID)
}
