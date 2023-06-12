package handler

import (
	"github.com/gin-gonic/gin"

	db "github.com/yegu-sanjana-ozone/mobile-app/pkg/mobile/DATABASE"
	service "github.com/yegu-sanjana-ozone/mobile-app/pkg/mobile/SERVICE"
)

func GetMobile ( c *gin.Context){
	cassandraSession := db.Session

	 mobile,err:= service.GetMobile(cassandraSession)
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, mobile)
}