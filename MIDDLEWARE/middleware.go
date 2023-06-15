package middleware

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"

	config "github.com/yegu-sanjana-ozone/mobile-app/cmd/configs"
	db "github.com/yegu-sanjana-ozone/mobile-app/pkg/mobile/DATABASE"
	Model "github.com/yegu-sanjana-ozone/mobile-app/pkg/mobile/MODEL"
	store "github.com/yegu-sanjana-ozone/mobile-app/pkg/mobile/STORE"
)



func RequireAuth(c *gin.Context) {
	cassandraSession := db.Session

	fmt.Println("bleh")

	tokenString := c.GetHeader("Authorization")

	if tokenString == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization token not found"})
		return
	}
	config, err := config.ReadConfig()
	fmt.Println(config)

	if err != nil {
		log.Fatal("Cannot load config", err.Error())
	}
	// validate it
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		fmt.Println("asdtest",token)
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); 
		!ok {
			fmt.Println("ok",ok )
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(config.SecretKey), nil
	})
	fmt.Println("token",token.Valid,token.Claims)
	


	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println("HERE")
		//check the expiry
		fmt.Println(claims,time.Now().Unix())
		if float64(time.Now().Unix()) > claims["expresat"].(float64) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token expired"})
			return
		}
		
		var user Model.User
        fmt.Println("claimsemail",claims["Email"].(string))
		user  = store.CheckEmail(cassandraSession, claims["Email"].(string))
  fmt.Println("user",user)
		fmt.Println("user",user)
		if user.Email == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"asdf": "Unauthorised"})
		}

		c.Set("user", user)
	c.Next()
	}
}
