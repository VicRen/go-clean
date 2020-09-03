package memory

import (
	"github.com/vicren/go-clean/domain/entity"
	"github.com/vicren/go-clean/domain/repository"
)

func NewUserRepository(users []entity.User) repository.UserRepository {
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
