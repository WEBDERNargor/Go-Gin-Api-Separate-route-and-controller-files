package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Usertype struct {
	Name string `json:"user_name"`
	Age  int32  `json:"user_age"`
}

var Usersdata = []Usertype{
	{
		Name: "Tom01",
		Age:  15,
	},
	{
		Name: "Tam",
		Age:  29,
	},
}

func GetUsers(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "succes",
		"data":   Usersdata,
	})
}

func PutUsers(c *gin.Context) {
	var user Usertype
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}
	Usersdata = append(Usersdata, user)
	c.JSON(200, gin.H{
		"status": "succes",
		"data":   Usersdata,
	})
}
