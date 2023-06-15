package Model

import "github.com/golang-jwt/jwt"

type Mobile struct {
	ID int `json:"ID"`
	Brand string `json:"brand"`
	Model string `json:"model"`
	Year  string `json:"year"`
	Price string `json:"price"`
}

type User struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

// type JwtCustomClaims struct {
// 	Email string `json:"name"`
// 	exp 
// }

type JwtCustomRefreshClaims struct {
	ID string `json:"id"`
	jwt.StandardClaims
}