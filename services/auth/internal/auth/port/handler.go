package port

import "github.com/gin-gonic/gin"

type ISignUpHandler interface {
	// (POST /v1/auth/register)
	SignUp(c *gin.Context)
	// (POST /v1/auth/admin/register)
	SignUpAdmin(c *gin.Context)
}

type ILoginHandler interface {
	// (POST /api/v1/auth/login)
	Login(c *gin.Context)
}

type IProfileHandler interface {
	// (GET /api/v1/auth/profile)
	GetProfile(c *gin.Context)
}
