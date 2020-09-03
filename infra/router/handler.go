package router

import (
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
	v1 "github.com/vicren/go-clean/infra/router/v1"
)

type Handler struct {
	config *Config
	router *gin.Engine
}

func (h *Handler) Start() error {
	if h.router != nil {
		return errors.New("engine already started")
	}
	h.router = gin.Default()

	// TODO create App controller

	switch h.config.APIVersion {
	case "v1":
		v1.Bind(h.router, nil)
	default:
		return fmt.Errorf("unsupported API verison: %s", h.config.APIVersion)
	}
	return h.router.Run(h.config.Address)
}

func (h *Handler) Close() error {
	return nil
}
