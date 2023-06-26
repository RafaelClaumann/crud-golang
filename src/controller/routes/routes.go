package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rafaelclaumann/crud-golang/src/controller"
)

func InitRoutes(r *gin.RouterGroup) {

	r.GET("/user/id/:userId", controller.GetUserById)
	r.GET("/user/email/:userEmail", controller.GetUserByEmail)
	r.POST("/user", controller.CreateUser)
	r.PUT("/user/:userId", controller.UpdateUser)
	r.DELETE("/user/:userId", controller.DeleteUser)

}
