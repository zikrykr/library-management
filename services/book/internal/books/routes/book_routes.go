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
	// create Book
	router.POST("", bookHandler.CreateBook)
	// update Book
	router.PUT("/:id", bookHandler.UpdateBook)
	// delete Book
	router.DELETE("/:id", bookHandler.DeleteBookByID)
	// get Books
	router.GET("", bookHandler.GetBooks)
	// get Books by id
	router.GET("/:id", bookHandler.GetBookByID)
}
