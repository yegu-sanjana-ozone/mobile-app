package users

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/yegu-sanjana-ozone/mobile-app/pkg/mobile/database"
	Model "github.com/yegu-sanjana-ozone/mobile-app/pkg/mobile/model"
	store "github.com/yegu-sanjana-ozone/mobile-app/pkg/mobile/store"
	tokenutil "github.com/yegu-sanjana-ozone/mobile-app/pkg/mobile/tokenutil"
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
		fmt.Println(token)

		if(err!=nil){
			panic(err)
		}
		c.JSON(http.StatusOK,gin.H{"token":token})
	}