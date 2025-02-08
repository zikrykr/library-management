package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/zikrykr/library-management/services/auth/internal/auth/port"
)

type (
	authPublicRoutes struct{}
	authRoutes       struct{}
	adminRoutes      struct{}
)

var (
	PublicRoutes authPublicRoutes
	Routes       authRoutes
)

func (r authPublicRoutes) NewPublicRoutes(router *gin.RouterGroup, signUpHandler port.ISignUpHandler, loginHandler port.ILoginHandler) {
	// sign-up
	router.POST("/register", signUpHandler.SignUp)
	// login
	router.POST("/login", loginHandler.Login)
}

func (r authRoutes) NewRoutes(router *gin.RouterGroup, profileHandler port.IProfileHandler) {
	// get profile
	router.GET("/me", profileHandler.GetProfile)
}

func (r authRoutes) NewAdminRoutes(router *gin.RouterGroup, signUpHandler port.ISignUpHandler) {
	rAdmin := router.Group("/admin")
	// register admin
	rAdmin.POST("/register", signUpHandler.SignUpAdmin)
}
