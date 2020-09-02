package main

import (
	"github.com/gin-gonic/gin"
	"github.com/vicren/go-clean/adapter/controller"
	"github.com/vicren/go-clean/domain/entity"
	"github.com/vicren/go-clean/infra/router"
	"github.com/vicren/go-clean/infra/storage"
)

func main() {
	r := gin.Default()

	s := &storage.MemoryStorage{
		Users: []*entity.User{
			{
				ID:    123,
				Name:  "testing",
				Email: "haha",
			},
		},
	}
	uc := controller.NewUserController(s)
	ac := &controller.AppController{
		User: uc,
	}

	router.NewRouter(r, ac)

	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}
