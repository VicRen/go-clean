package model

import "github.com/vicren/go-clean/domain/entity"

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func ParseUser(user entity.User) *User {
	return &User{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}
}

func ParseUsers(users []entity.User) []*User {
	var ret []*User
	for _, u := range users {
		ret = append(ret, ParseUser(u))
	}
	return ret
}
