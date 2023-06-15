package handler

import (
	"github.com/gin-gonic/gin"
	Model "github.com/yegu-sanjana-ozone/mobile-app/pkg/mobile/MODEL"
	"net/http"

)

// CreateMobiles godoc
// @Summary Create a mobile
// @Description Create a new mobile with the provided details
// @Tags Mobile
// @Accept json
// @Produce json
// @Param mobile body MobileRequest true "Mobile object to be created"
// @Success 200 {string} string "Mobile created successfully"
// @Failure 400 {object} ErrorResponse "Error response"
// @Router /mobile [post]

func (h *Handler) CreateMobiles (c *gin.Context) {

 var request Model.Mobile

 c.ShouldBindJSON(&request)

 mobile := Model.Mobile{
	ID:request.ID,
	Brand: request.Brand,
	Model: request.Model,
	Year: request.Year,
    Price: request.Price,
 }

 err := h.service.CreateMobile(mobile)

 if err!= nil {
     c.JSON(400, gin.H{"error": err.Error()})
     return
 }

c.JSON(http.StatusOK,gin.H{"message": "Created"})
}
