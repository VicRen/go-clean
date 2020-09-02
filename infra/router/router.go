package router

import (
	"github.com/gin-gonic/gin"
	"github.com/vicren/go-clean/adapter/controller"
)

func NewRouter(router *gin.Engine, c *controller.AppController) {

	router.GET("/users", func(context *gin.Context) {
		c.User.GetUsers(context)
	})
}
