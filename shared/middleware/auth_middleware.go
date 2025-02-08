package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/zikrykr/library-management/shared/constants"
	"github.com/zikrykr/library-management/shared/pkg"
)

func ParseTokenFromHeader(c *gin.Context) (string, error) {
	var (
		headerToken = c.Request.Header.Get("Authorization")
		splitToken  []string
	)

	splitToken = strings.Split(headerToken, "Bearer ")

	// check valid bearer token
	if len(splitToken) <= 1 {
		return "", errors.New("invalid token")
	}

	return splitToken[1], nil
}

func JwtAuthMiddleware(jwtSecret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// token claims
		headerToken, err := ParseTokenFromHeader(c)
		if err != nil {
			pkg.ResponseError(c, http.StatusUnauthorized, err)
			return
		}

		claims := &pkg.JWTClaims{}
		token, err := jwt.ParseWithClaims(headerToken, claims, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok { // check signing method
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(jwtSecret), nil
		})
		// check parse token error
		if err != nil || !token.Valid {
			pkg.ResponseError(c, http.StatusUnauthorized, err)
			return
		}

		c.Set(constants.CONTEXT_CLAIM_USER_EMAIL, claims.Email)
		c.Set(constants.CONTEXT_CLAIM_USER_ID, claims.UserID)
		c.Set(constants.CONTEXT_CLAIM_USER_ROLE, claims.Role)
		c.Set(constants.CONTEXT_CLAIM_KEY, claims)

		c.Next()
	}
}
