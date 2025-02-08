package middleware

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/zikrykr/library-management/shared/constants"
	"github.com/zikrykr/library-management/shared/pkg"
)

func CheckAdminRole() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get(constants.CONTEXT_CLAIM_USER_ROLE)
		if !exists {
			pkg.ResponseError(c, http.StatusUnauthorized, errors.New("role not found in context"))
			return
		}

		// Convert role to string
		roleStr, ok := role.(string)
		if !ok {
			pkg.ResponseError(c, http.StatusUnauthorized, errors.New("role type assertion failed"))
			return
		}

		if !strings.EqualFold(roleStr, constants.ROLE_ADMIN) {
			pkg.ResponseError(c, http.StatusUnauthorized, errors.New("Unauthorized, only admin can access this endpoint"))
		}

		c.Next()
	}
}
