package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rafaelclaumann/crud-golang/src/configuration/rest_error"
	"github.com/rafaelclaumann/crud-golang/src/model/request"
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
	var requestBody request.UserRequest
	err := c.ShouldBindJSON(&requestBody)

	if err != nil {
		err_msg := fmt.Sprintf("[CreateUser] there are some incorrect fields in request body! err [ %s ]", err.Error())
		restErr := rest_error.NewBadRequestError(err_msg)

		c.JSON(http.StatusBadRequest, restErr)
		return
	}

	c.JSON(http.StatusOK, requestBody)
}

func UpdateUser(c *gin.Context) {
	err := rest_error.NewInternalServerError("[UpdateUser] not implemented yet")
	c.JSON(http.StatusInternalServerError, err)
}

func DeleteUser(c *gin.Context) {
	err := rest_error.NewInternalServerError("[DeleteUser] not implemented yet")
	c.JSON(http.StatusInternalServerError, err)
}
