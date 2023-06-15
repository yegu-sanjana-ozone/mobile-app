package handler

import (
	"strconv"

	"github.com/gin-gonic/gin"
	Model "github.com/yegu-sanjana-ozone/mobile-app/pkg/mobile/MODEL"

)

func (h *Handler) UpdateByID(c *gin.Context) {
	
	str := c.Param("id")
	id , _ := strconv.Atoi(str)
	var request Model.Mobile

 c.ShouldBindJSON(&request)
 
	mob := h.service.UpdateByID(id,request.Brand)
	c.JSON(200,mob)
}