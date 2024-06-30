package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Producttype struct {
	Name  string `json:"product_name"`
	Price int32  `json:"product_price"`
}

var Productsdata = []Producttype{
	{
		Name:  "Comset01",
		Price: 5000,
	},
	{
		Name:  "Comset02",
		Price: 6000,
	},
}

func GetProducts(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "succes",
		"data":   Productsdata,
	})
}

func PutProducts(c *gin.Context) {
	var product Producttype
	if err := c.BindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	Productsdata = append(Productsdata, product)
	c.JSON(200, gin.H{
		"status": "succes",
		"data":   Productsdata,
	})
}
