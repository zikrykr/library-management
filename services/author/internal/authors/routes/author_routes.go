package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/zikrykr/library-management/services/author/internal/authors/port"
)

type (
	authorRoutes struct{}
)

var (
	Routes authorRoutes
)

func (r authorRoutes) NewRoutes(router *gin.RouterGroup, authorHandler port.IAuthorHandler) {
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
