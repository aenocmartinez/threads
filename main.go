package main

import (
	"threads/src/infraestructure/middleware"
	"threads/src/view/controller"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.POST("/login", controller.Login)
	r.POST("/register", controller.Register)
	r.GET("/check-db", controller.CheckDBConnection)

	protected := r.Group("/api")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.POST("/logout", controller.Logout)
		// protected.GET("/test", controller.TestEndpoint)
	}

	// Puerto de ejecución
	r.Run(":8590")
}
