package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/zikrykr/library-management/services/author/internal/authors/port"
)

type (
	routes      struct{}
	adminRoutes struct{}
)

var (
	Routes      routes
	AdminRoutes adminRoutes
)

func (r adminRoutes) NewAdminRoutes(router *gin.RouterGroup, authorHandler port.IAuthorHandler) {
	// get authors
	router.GET("", authorHandler.GetAuthors)
	// get authors by id
	router.GET("/:id", authorHandler.GetAuthorByID)
	// create author
	router.POST("", authorHandler.CreateAuthor)
	// update author
	router.PUT("/:id", authorHandler.UpdateAuthor)
	// delete author
	router.DELETE("/:id", authorHandler.DeleteAuthorByID)
}

func (r routes) NewRoutes(router *gin.RouterGroup, authorHandler port.IAuthorHandler) {
	// get authors
	router.GET("", authorHandler.GetAuthors)
	// get authors by id
	router.GET("/:id", authorHandler.GetAuthorByID)
}
