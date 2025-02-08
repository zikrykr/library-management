package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/zikrykr/library-management/services/book/internal/books/port"
)

type (
	bookAdminRoutes struct{}
	bookRoutes      struct{}
)

var (
	AdminRoutes bookAdminRoutes
	Routes      bookRoutes
)

func (r bookRoutes) NewRoutes(router *gin.RouterGroup, bookHandler port.IBookHandler) {
	// get Books
	router.GET("", bookHandler.GetBooks)
	// get Books by id
	router.GET("/:id", bookHandler.GetBookByID)
}

func (r bookAdminRoutes) NewAdminRoutes(router *gin.RouterGroup, bookHandler port.IBookHandler) {
	rAdmin := router.Group("/admin")
	// create Book
	rAdmin.POST("", bookHandler.CreateBook)
	// update Book
	rAdmin.PUT("/:id", bookHandler.UpdateBook)
	// delete Book
	rAdmin.DELETE("/:id", bookHandler.DeleteBookByID)
	// get Books
	rAdmin.GET("", bookHandler.GetBooks)
	// get Books by id
	rAdmin.GET("/:id", bookHandler.GetBookByID)
}
