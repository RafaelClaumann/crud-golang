package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rafaelclaumann/crud-golang/src/configuration/rest_error"
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
	err := rest_error.NewInternalServerError("[UpdateUser] not implemented yet")
	c.JSON(http.StatusInternalServerError, err)
}

func DeleteUser(c *gin.Context) {
	err := rest_error.NewInternalServerError("[DeleteUser] not implemented yet")
	c.JSON(http.StatusInternalServerError, err)
}
