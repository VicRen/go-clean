package repository

import "github.com/vicren/go-clean/domain/entity"

type UserRepository interface {
	FindAll() ([]entity.User, error)
}