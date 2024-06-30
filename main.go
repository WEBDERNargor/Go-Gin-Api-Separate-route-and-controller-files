package main

import (
	"go-api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	route := gin.Default()
	routes.ProductRoute(route)
	routes.UserRoute(route)
	route.Run("0.0.0.0:3001")
}
