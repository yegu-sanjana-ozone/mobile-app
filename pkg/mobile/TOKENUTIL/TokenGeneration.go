package tokenutil

import (
	"time"

	"github.com/golang-jwt/jwt"
	Model "github.com/yegu-sanjana-ozone/mobile-app/pkg/mobile/MODEL"
)

func CreateAccessToken(user *Model.User, secret string, expiry int) (accessToken string, err error) {
	
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,jwt.MapClaims {
		"Email": user.Email,
		"expresat": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})
	t, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return t, err
}
