package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/yegu-sanjana-ozone/mobile-app/pkg/mobile/DATABASE"
	Model "github.com/yegu-sanjana-ozone/mobile-app/pkg/mobile/MODEL"
	store "github.com/yegu-sanjana-ozone/mobile-app/pkg/mobile/STORE"
	tokenutil "github.com/yegu-sanjana-ozone/mobile-app/pkg/mobile/TOKENUTIL"
)


func Login(c *gin.Context) {
	cassandraSession := db.Session
	var user Model.User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(400, gin.H{"error": "Invalid request payload"})
			return
		}

		err := store.ValidateUser(cassandraSession, user)
		if err != nil {
			c.JSON(401, gin.H{"error": "Invalid email or password"})
			return
		}

		c.JSON(200, gin.H{"message": "Login successful"})

		token,err := tokenutil.CreateAccessToken(&user,"hello",913208515142266763)

		if(err!=nil){
			panic(err)
		}
		c.JSON(http.StatusOK,gin.H{"token":token})
	}