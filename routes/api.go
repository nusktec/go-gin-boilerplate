package routes

import (
	"goginapi/controllers"
	"goginapi/middleware"
	"os"

	"github.com/gin-gonic/gin"
)

type ApiRoutes struct{}

func (api ApiRoutes) GetApiRoutes(router *gin.Engine) *gin.Engine {

	// controller
	user := new(controllers.UserController)

	// routes
	v1 := router.Group("/api/v1")
	v1.Use(middleware.AuthMiddleware())

	v1.GET("/", user.Retrieve)
	v1.GET("/user/:id", user.GetSingle)

	v1.GET("env", func(c *gin.Context) {

		s3Bucket := os.Getenv("SECRET_KEY")

		c.JSON(200, gin.H{
			"s3-bucket": s3Bucket,
		})
	})

	return router

}
