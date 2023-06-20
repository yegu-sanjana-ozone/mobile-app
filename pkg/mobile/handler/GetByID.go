package handler

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetByID(c *gin.Context){

	str := c.Param("id")
	id,_ := strconv.Atoi(str)
	mob := h.service.GetByID(id)
    c.JSON(200,mob)
}