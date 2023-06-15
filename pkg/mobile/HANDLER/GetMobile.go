package handler

import (
	"github.com/gin-gonic/gin"

)

func (h *Handler) GetMobile ( c *gin.Context){


	 mobile,err:= h.service.GetMobile()
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, mobile)
}