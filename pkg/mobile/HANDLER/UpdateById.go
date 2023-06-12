package handler

import (
	"strconv"

	"github.com/gin-gonic/gin"
	db "github.com/yegu-sanjana-ozone/mobile-app/pkg/mobile/DATABASE"
	Model "github.com/yegu-sanjana-ozone/mobile-app/pkg/mobile/MODEL"
	service "github.com/yegu-sanjana-ozone/mobile-app/pkg/mobile/SERVICE"
)

func UpdateByID(c *gin.Context) {
	cassandraSession := db.Session
	str := c.Param("id")
	id , _ := strconv.Atoi(str)
	var request Model.Mobile

 c.ShouldBindJSON(&request)
 
	mob := service.UpdateByID(id,cassandraSession,request.Brand)
	c.JSON(200,mob)
}