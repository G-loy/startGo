package user

import (
	"firstGin/models/request"
	"github.com/gin-gonic/gin"
	"net/http"
)

func userLogin(c *gin.Context) {
	json := request.User{}
	c.Bind(&json)
	c.JSON(http.StatusOK, gin.H{
		"message":  "POST",
		"name":     json.Name,
		"password": json.Password,
	})
}
