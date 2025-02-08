package port

import "github.com/gin-gonic/gin"

type IAuthorHandler interface {
	// (POST /v1/authors)
	CreateAuthor(c *gin.Context)
	// (PUT /v1/authors/{id})
	UpdateAuthor(c *gin.Context)
	// (GET /v1/authors)
	GetAuthors(c *gin.Context)
	// (POST /v1/authors/{id})
	GetAuthorByID(c *gin.Context)
	// (DELETE /v1/authors/{id})
	DeleteAuthorByID(c *gin.Context)
}
