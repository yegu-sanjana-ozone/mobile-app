package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/yegu-sanjana-ozone/mobile-app/pkg/mobile/DATABASE"
	Model "github.com/yegu-sanjana-ozone/mobile-app/pkg/mobile/MODEL"
	store "github.com/yegu-sanjana-ozone/mobile-app/pkg/mobile/STORE"
	"golang.org/x/crypto/bcrypt"
)

func Signup(c *gin.Context) {
	//Get email and pass from the body
	cassandraSession := db.Session
	var body struct {
		Email    string
		Password string
	}
	err := c.Bind(&body)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Payload"})
		return
	}

	//Hash the password
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//Create the user
	user := Model.User{Email: body.Email, Password: string(hash)}

	store.CreateUser(cassandraSession, user )
	//Send a response
	c.JSON(http.StatusOK, gin.H{"data": "Created successfully"})

}
