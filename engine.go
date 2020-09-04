package core

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/vicren/go-clean/adapter/controller"
	"github.com/vicren/go-clean/infra/router"
)

type Server interface {
}

type Instance struct {
	config *Config
	router *gin.Engine
}

func (h *Instance) Start() error {
	if h.router != nil {
		return errors.New("engine already started")
	}
	if len(h.config.APIVersions) == 0 {
		return errors.New("API version not specified")
	}
	h.router = gin.Default()

	// TODO create App controller
	c := &controller.AppController{}

	for _, v := range h.config.APIVersions {
		rf, err := router.GetBindFunc(v)
		if err != nil {
			return fmt.Errorf("unsupported API verison: %s>%v", v, err)
		}
		rf(h.router, c)
	}

	return h.router.Run(h.config.Address)
}

func (h *Instance) Close() error {
	return nil
}

type resolution struct {
	deps     []reflect.Type
	callback interface{}
}
