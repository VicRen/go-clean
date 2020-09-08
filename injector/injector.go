// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package injector

import (
	"github.com/google/wire"
	"github.com/vicren/go-clean/adapter/controller"
	"github.com/vicren/go-clean/adapter/repository/memory"
	"github.com/vicren/go-clean/domain/entity"
	"github.com/vicren/go-clean/domain/interactor"
)

var MemoryStorageSet = wire.NewSet(controller.ProvideAppController, controller.UserSet, interactor.UserSet, memory.RepoSet)

func InitMemoryEngine(Users []entity.User) *controller.AppController {
	wire.Build(MemoryStorageSet)
	return nil
}
