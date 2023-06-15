package handler

import (
	"strconv"

	"github.com/gin-gonic/gin"

)

// DeleteMobile godoc
// @Summary Delete a mobile
// @Description Deletes a mobile with the specified ID
// @Tags Mobile
// @Accept       json
// @Produce      json
// @Param id path int true "Mobile ID"
// @Success 200 {object} object
// @Router /mobile/{id} [delete]

func (h *Handler) DeleteByID(c *gin.Context){
	str := c.Param("id")
	id,_ := strconv.Atoi(str)
	mob := h.service.DeleteMobile(id)
	c.JSON(200,mob)
}