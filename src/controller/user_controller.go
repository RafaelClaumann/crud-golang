package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUserById(c *gin.Context) {
	userId := c.Param("userId")

	c.JSON(http.StatusOK, gin.H{
		"GetUserById": userId,
	})
}

func GetUserByEmail(c *gin.Context) {
	userEmail := c.Param("userEmail")

	c.JSON(http.StatusOK, gin.H{
		"GetUserByEmail": userEmail,
	})
}

func CreateUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"CreateUser": "create-user",
	})
}

func UpdateUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"UpdateUser": "update-user",
	})
}

func DeleteUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"DeleteUser": "delete-user",
	})
}
