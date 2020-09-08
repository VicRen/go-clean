package memory

import (
	"github.com/vicren/go-clean/domain/entity"
)

func NewUserRepository(users []entity.User) *UserRepository {
	return &UserRepository{
		Users: users,
	}
}

type UserRepository struct {
	Users []entity.User
}

func (r *UserRepository) FindAll() ([]entity.User, error) {
	return r.Users, nil
}
