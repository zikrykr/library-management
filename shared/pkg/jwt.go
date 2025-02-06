package pkg

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

const (
	JWT_SUBJECT = "access_token"
)

type JWTClaims struct {
	jwt.RegisteredClaims
	UserID string `json:"user_id"`
	Email  string `json:"email"`
}

type JWTResp struct {
	AccessToken string `json:"access_token"`
}

type JWTConfig struct {
	AppName   string
	JWTSecret string
}

func GenerateJWT(claimsData *JWTClaims, config JWTConfig) (JWTResp, error) {
	issuedAt := time.Now()
	expAt := issuedAt.Add(60 * time.Minute)

	claimsData.RegisteredClaims = jwt.RegisteredClaims{
		Issuer:    config.AppName,
		Subject:   JWT_SUBJECT,
		ExpiresAt: jwt.NewNumericDate(expAt),
		IssuedAt:  jwt.NewNumericDate(issuedAt),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, *claimsData)

	signedToken, err := token.SignedString([]byte(config.JWTSecret))
	if err != nil {
		return JWTResp{}, err
	}

	return JWTResp{
		AccessToken: signedToken,
	}, nil
}
