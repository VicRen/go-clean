// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package injector

import (
	"github.com/google/wire"
	"github.com/vicren/go-clean/adapter/controller"
	"github.com/vicren/go-clean/adapter/repository/memory"
	"github.com/vicren/go-clean/domain/entity"
	"github.com/vicren/go-clean/domain/interactor"
)

// Injectors from injector.go:

func InitMemoryEngine(Users []entity.User) *controller.AppController {
	userRepository := memory.NewUserRepository(Users)
	userController := controller.NewUserController(userRepository)
	appController := controller.ProvideAppController(userController)
	return appController
}

// injector.go:

var MemoryStorageSet = wire.NewSet(controller.ProvideAppController, controller.UserSet, interactor.UserSet, memory.RepoSet)
