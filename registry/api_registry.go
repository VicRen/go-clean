package registry

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/vicren/go-clean/adapter/controller"
)

type BindFunc func(router *gin.Engine, c *controller.AppController)

var bindFuncRegistry = make(map[string]BindFunc)

func RegisterBindFunc(version string, bindFunc BindFunc) error {
	if _, found := bindFuncRegistry[version]; found {
		return fmt.Errorf("bind function for verison %s is already registered", version)
	}
	bindFuncRegistry[version] = bindFunc
	return nil
}
