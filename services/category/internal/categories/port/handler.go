package port

import "github.com/gin-gonic/gin"

type ICategoryHandler interface {
	// (POST /v1/categories)
	CreateCategory(c *gin.Context)
	// (PUT /v1/categories/{id})
	UpdateCategory(c *gin.Context)
	// (GET /v1/categories)
	GetCategories(c *gin.Context)
	// (POST /v1/categories/{id})
	GetCategoryByID(c *gin.Context)
	// (DELETE /v1/categories/{id})
	DeleteCategoryByID(c *gin.Context)
}
