package controller

type Context interface {
	JSON(code int, obj interface{})
}
