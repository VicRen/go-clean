package repository

import (
	"github.com/vicren/go-clean/domain/entity"
	"github.com/vicren/go-clean/domain/repository"
)

func NewUserRepository(storage Storage) repository.UserRepository {
	return &userRepository{storage: storage}
}

type userRepository struct {
	storage Storage
}

func (r *userRepository) FindAll() ([]entity.User, error) {
	var users []*entity.User
	if err := r.storage.Find(&users); err != nil {
		return nil, err
	}
	var ret []entity.User
	for _, u := range users {
		ret = append(ret, *u)
	}
	return ret, nil
}
