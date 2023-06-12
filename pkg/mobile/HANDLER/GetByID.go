package handler

import (
	"strconv"

	"github.com/gin-gonic/gin"
	db "github.com/yegu-sanjana-ozone/mobile-app/pkg/mobile/DATABASE"
	service "github.com/yegu-sanjana-ozone/mobile-app/pkg/mobile/SERVICE"
)

func GetByID(c *gin.Context){
	cassandraSession := db.Session
	str := c.Param("id")
	id,_ := strconv.Atoi(str)
	mob := service.GetByID(id,cassandraSession)
    c.JSON(200,mob)
}