package tokenutil

import (
	"time"

	"github.com/golang-jwt/jwt"
	Model "github.com/yegu-sanjana-ozone/mobile-app/pkg/mobile/MODEL"
)

func CreateAccessToken(user *Model.User, secret string, expiry int) (accessToken string, err error) {
	exp := time.Now().Add(time.Hour * time.Duration(expiry)).Unix()
	claims := &Model.JwtCustomClaims{
		Email: user.Email,
		ID:   "1234",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: exp,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return t, err
}
