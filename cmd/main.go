package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	middleware "github.com/yegu-sanjana-ozone/mobile-app/MIDDLEWARE"
	db "github.com/yegu-sanjana-ozone/mobile-app/pkg/mobile/DATABASE"
	handler "github.com/yegu-sanjana-ozone/mobile-app/pkg/mobile/HANDLER"
	users "github.com/yegu-sanjana-ozone/mobile-app/pkg/mobile/USER"
)



func InitRoutes(adminRoutes *gin.RouterGroup){
    mh := handler.New()
	mobileRoute := adminRoutes.Group("/mobile")
	baseRoute := adminRoutes.Group("")
	baseRouteAuth := adminRoutes.Group("")
	mobileRoute.Use(middleware.RequireAuth)
	baseRouteAuth.Use(middleware.RequireAuth)
	mobileRoute.POST("",mh.CreateMobiles)
	mobileRoute.GET("", mh.GetMobile)
	mobileRoute.GET("/:id",mh.GetByID)
	mobileRoute.DELETE("/:id",mh.DeleteByID)
    mobileRoute.PUT("/:id",mh.UpdateByID)
	baseRoute.POST("/signup",users.Signup)
	baseRoute.POST("/login",users.Login)
	baseRouteAuth.GET("/validate",users.Validate)
}

func main() {
	ses:=db.Session
	fmt.Println("ses",ses)
	r := gin.Default()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	adminRoutes := r.Group("")
	InitRoutes(adminRoutes)
	r.Run()
}