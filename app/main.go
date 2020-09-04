package main

import (
	"github.com/gin-gonic/gin"
	"github.com/vicren/go-clean/adapter/controller"
	"github.com/vicren/go-clean/adapter/repository/memory"
	"github.com/vicren/go-clean/domain/entity"
	v1 "github.com/vicren/go-clean/infra/router/v1"
)

func main() {
	r := gin.Default()

	uc := controller.NewUserController(memory.NewUserRepository([]entity.User{
		{
			ID:    1,
			Name:  "testing",
			Email: "testing@testing.com",
		},
		{
			ID:    2,
			Name:  "testing2",
			Email: "testing2@testing.com",
		},
	}))
	ac := &controller.AppController{
		User: uc,
	}

	v1.Bind(r, ac)

	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}
