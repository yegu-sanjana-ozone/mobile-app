package handler

import (
	"github.com/gin-gonic/gin"
	Model "github.com/yegu-sanjana-ozone/mobile-app/pkg/mobile/MODEL"
	service "github.com/yegu-sanjana-ozone/mobile-app/pkg/mobile/SERVICE"
)

func CreateMobiles (c *gin.Context) {

 var request Model.Mobile

 c.ShouldBindJSON(&request)

 mobile := Model.Mobile{
	ID:request.ID,
	Brand: request.Brand,
	Model: request.Model,
	Year: request.Year,
    Price: request.Price,
 }

 err := service.CreateMobile(mobile)

 if err!= nil {
     c.JSON(400, gin.H{"error": err.Error()})
     return
 }


}