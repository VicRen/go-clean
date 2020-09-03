package controller

import (
	"reflect"
	"testing"

	"github.com/vicren/go-clean/domain/entity"
)

func TestParseUser(t *testing.T) {
	type args struct {
		user entity.User
	}
	tests := []struct {
		name string
		args args
		want *RespUser
	}{
		{
			"happy_case",
			args{
				entity.User{
					ID:    1,
					Name:  "testing",
					Email: "testing@testing.com",
				},
			},
			&RespUser{
				ID:    1,
				Name:  "testing",
				Email: "testing@testing.com",
			},
		},
		{
			"happy_case_no_email",
			args{
				entity.User{
					ID:   1,
					Name: "testing",
				},
			},
			&RespUser{
				ID:   1,
				Name: "testing",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseUser(tt.args.user); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseUsers(t *testing.T) {
	type args struct {
		users []entity.User
	}
	tests := []struct {
		name string
		args args
		want []*RespUser
	}{
		{
			"happy_case",
			args{
				[]entity.User{
					{
						ID:    1,
						Name:  "testing",
						Email: "testing@testing.com",
					},
				},
			},
			[]*RespUser{
				{
					ID:    1,
					Name:  "testing",
					Email: "testing@testing.com",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseUsers(tt.args.users); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseUsers() = %v, want %v", got, tt.want)
			}
		})
	}
}
