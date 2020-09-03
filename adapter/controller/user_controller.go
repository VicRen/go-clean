package controller

import (
	"net/http"

	"github.com/vicren/go-clean/domain/entity"

	"github.com/vicren/go-clean/domain/interactor"
	"github.com/vicren/go-clean/domain/repository"
)

type UserController interface {
	GetUsers(c Context)
}

func NewUserController(userRepository repository.UserRepository) UserController {
	return &userController{
		userInteractor: interactor.NewUserInteractor(userRepository),
	}
}

type userController struct {
	userInteractor interactor.UserInteractor
}

func (uc *userController) GetUsers(c Context) {
	type (
		Response struct {
			Code  int         `json:"code"`
			Users []*RespUser `json:"users"`
		}
	)
	u, err := uc.userInteractor.List()
	if err != nil {
		e := NewError(http.StatusInternalServerError, err.Error())
		c.JSON(e.Code, e)
		return
	}

	resp := &Response{
		Code:  http.StatusOK,
		Users: ParseUsers(u),
	}

	c.JSON(resp.Code, resp)
}

type RespUser struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func ParseUser(user entity.User) *RespUser {
	return &RespUser{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}
}

func ParseUsers(users []entity.User) []*RespUser {
	var ret []*RespUser
	for _, u := range users {
		ret = append(ret, ParseUser(u))
	}
	return ret
}
