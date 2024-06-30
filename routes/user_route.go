package routes

import (
	"go-api/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {

	productGroup := router.Group("/user")
	{
		productGroup.GET("/", controllers.GetUsers)
		productGroup.PUT("/", controllers.PutUsers)
	}

}
