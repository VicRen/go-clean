package main

import (
	"github.com/gin-gonic/gin"
	"github.com/vicren/go-clean/domain/entity"
	v1 "github.com/vicren/go-clean/infra/router/v1"
	"github.com/vicren/go-clean/injector"
)

func main() {
	r := gin.Default()

	users := []entity.User{
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
		{
			ID:    3,
			Name:  "testing3",
			Email: "testing3@testing.com",
		},
	}

	//uc := controller.NewUserController(memory.NewUserRepository(users))
	//ac := &controller.AppController{
	//	User: uc,
	//}

	ac := injector.InitMemoryEngine(users)

	v1.Bind(r, ac)

	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}
