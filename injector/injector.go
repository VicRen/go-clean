package registry

import (
	"github.com/google/wire"
	"github.com/vicren/go-clean/adapter/controller"
	"github.com/vicren/go-clean/adapter/repository/memory"
	"github.com/vicren/go-clean/domain/entity"
	"github.com/vicren/go-clean/domain/interactor"
)

var MemoryStorageSet = wire.NewSet(interactor.NewUserInteractor, memory.NewUserRepository)

func InitMemoryEngine(Users []entity.User) *controller.AppController {
	wire.Build(MemoryStorageSet)
	return nil
}
