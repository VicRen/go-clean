package interactor

import (
	"errors"

	"github.com/vicren/go-clean/domain/entity"
	"github.com/vicren/go-clean/domain/repository"
)

type UserInteractor interface {
	List() ([]entity.User, error)
}

func NewUserInteractor(userRepository repository.UserRepository) *UserUseCase {
	return &UserUseCase{
		userRepository: userRepository,
	}
}

type UserUseCase struct {
	userRepository repository.UserRepository
}

func (i *UserUseCase) List() ([]entity.User, error) {
	if i.userRepository == nil {
		return nil, errors.New("invalid user repository")
	}
	return i.userRepository.FindAll()
}
