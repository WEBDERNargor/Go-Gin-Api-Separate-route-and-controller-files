package routes

import (
	"go-api/controllers"

	"github.com/gin-gonic/gin"
)

func ProductRoute(router *gin.Engine) {

	productGroup := router.Group("/product")
	{
		productGroup.GET("/", controllers.GetProducts)
		productGroup.PUT("/", controllers.PutProducts)
	}

}
