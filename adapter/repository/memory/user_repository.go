package memory

import (
	"github.com/google/wire"
	"github.com/vicren/go-clean/domain/entity"
	"github.com/vicren/go-clean/domain/repository"
)

func NewUserRepository(users []entity.User) *userRepository {
	return &userRepository{
		Users: users,
	}
}

type userRepository struct {
	Users []entity.User
}

func (r *userRepository) FindAll() ([]entity.User, error) {
	return r.Users, nil
}

var Provider = wire.NewSet(NewUserRepository, wire.Bind(new(repository.UserRepository), new(*userRepository)))
