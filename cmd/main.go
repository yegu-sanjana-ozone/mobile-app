package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	db "github.com/yegu-sanjana-ozone/mobile-app/pkg/mobile/DATABASE"
	handler "github.com/yegu-sanjana-ozone/mobile-app/pkg/mobile/HANDLER"
	users "github.com/yegu-sanjana-ozone/mobile-app/pkg/mobile/USER"
)



func InitRoutes(adminRoutes *gin.RouterGroup){
	mobileRoute := adminRoutes.Group("/mobile")
	baseRoute := adminRoutes.Group("")
	mobileRoute.POST("",handler.CreateMobiles)
	mobileRoute.GET("", handler.GetMobile)
	mobileRoute.GET("/:id",handler.GetByID)
	mobileRoute.DELETE("/:id",handler.DeleteByID)
    mobileRoute.PUT("/:id",handler.UpdateByID)
	baseRoute.POST("/signup",users.Signup)
	baseRoute.POST("/login",users.Login)
}

func main() {
	ses:=db.Session
	fmt.Println("ses",ses)
	r := gin.Default()
	adminRoutes := r.Group("")
	InitRoutes(adminRoutes)
	r.Run()
}