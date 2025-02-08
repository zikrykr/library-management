package port

import "github.com/gin-gonic/gin"

type IBookHandler interface {
	// (POST /v1/books)
	CreateBook(c *gin.Context)
	// (PUT /v1/books/{id})
	UpdateBook(c *gin.Context)
	// (GET /v1/books)
	GetBooks(c *gin.Context)
	// (POST /v1/books/{id})
	GetBookByID(c *gin.Context)
	// (DELETE /v1/books/{id})
	DeleteBookByID(c *gin.Context)
}
