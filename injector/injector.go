// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package injector

import (
	"github.com/google/wire"
	"github.com/vicren/go-clean/adapter/controller"
	"github.com/vicren/go-clean/adapter/repository/memory"
	"github.com/vicren/go-clean/domain/entity"
)

var MemoryProvider = wire.NewSet(controller.NewAppController, controller.UserProvider, memory.Provider)

func InitMemoryEngine(Users []entity.User) *controller.AppController {
	wire.Build(MemoryProvider)
	return nil
}
