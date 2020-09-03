package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/vicren/go-clean/adapter/controller"
	"github.com/vicren/go-clean/common"
	"github.com/vicren/go-clean/registry"
)

func Bind(router *gin.Engine, c *controller.AppController) {

	router.GET("/v1/users", func(context *gin.Context) {
		c.User.GetUsers(context)
	})
}

func init() {
	common.Must(registry.RegisterBindFunc("v1", Bind))
}
