package controller

import (
	"net/http"

	"github.com/vicren/go-clean/adapter/model"
	"github.com/vicren/go-clean/adapter/repository"
	"github.com/vicren/go-clean/domain/interactor"
)

type UserController interface {
	GetUsers(c Context)
}

func NewUserController(storage repository.Storage) UserController {
	ur := repository.NewUserRepository(storage)
	return &userController{
		userInteractor: interactor.NewUserInteractor(ur),
	}
}

type userController struct {
	userInteractor interactor.UserInteractor
}

func (uc *userController) GetUsers(c Context) {
	u, err := uc.userInteractor.List()
	if err != nil {
		c.JSON(http.StatusInternalServerError, NewError(http.StatusInternalServerError, err.Error()))
		return
	}

	c.JSON(http.StatusOK, model.ParseUsers(u))
}
