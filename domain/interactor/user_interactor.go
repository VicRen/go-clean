package interactor

import (
	"errors"

	"github.com/google/wire"
	"github.com/vicren/go-clean/domain/entity"
	"github.com/vicren/go-clean/domain/repository"
)

type UserInteractor interface {
	List() ([]entity.User, error)
}

func NewUserInteractor(userRepository repository.UserRepository) *userInteractor {
	return &userInteractor{
		userRepository: userRepository,
	}
}

type userInteractor struct {
	userRepository repository.UserRepository
}

func (i *userInteractor) List() ([]entity.User, error) {
	if i.userRepository == nil {
		return nil, errors.New("invalid user repository")
	}
	return i.userRepository.FindAll()
}

var UserSet = wire.NewSet(NewUserInteractor, wire.Bind(new(UserInteractor), new(*userInteractor)))
