package routes

import (
	"github.com/RichardtJustke/crudGoooo/src/controller"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup) {
	r.GET("/getUserById/:userId", controller.FindUserById)
	r.GET("/getUserByEmail/:userEmail", controller.FindUserByEmail)
	r.POST("/createUser", controller.CreateUser)
	r.PUT("/updateUser/:userId", controller.UdpateUser)
	r.DELETE("/deleteUser/:userId", controller.DeleteUser)
}
