package interactor

import (
	"errors"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/vicren/go-clean/domain/entity"
	mock "github.com/vicren/go-clean/domain/repository/mocks"
)

var (
	mockUserRepository *mock.MockUserRepository
	userInstance       UserInteractor
)

func setupTestUserInteractor(t *testing.T) func(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockUserRepository = mock.NewMockUserRepository(ctrl)
	userInstance = NewUserInteractor(mockUserRepository)
	return func(t *testing.T) {
		mockUserRepository = nil
		ctrl.Finish()
	}
}

func Test_userInteractor_List_Nil_UserRepository(t *testing.T) {
	u := NewUserInteractor(nil)
	_, err := u.List()
	if !reflect.DeepEqual(err, errors.New("invalid user repository")) {
		t.Errorf("userInteractor.List() error = %v", err)
	}
}

func Test_userInteractor_List(t *testing.T) {
	tt := []struct {
		name    string
		want    []entity.User
		wantErr error
	}{
		{
			"happy_case",
			[]entity.User{
				{
					ID:    1,
					Name:  "user",
					Email: "user@user.com",
				},
			},
			nil,
		},
		{
			"error_case",
			nil,
			errors.New("error finding users"),
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			tearDown := setupTestUserInteractor(t)
			mockUserRepository.EXPECT().FindAll().Return(tc.want, tc.wantErr)

			got, err := userInstance.List()
			if !reflect.DeepEqual(err, tc.wantErr) {
				t.Errorf("userInteractor.List() error = %v, wantErr %v", err, tc.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("userInteractor.List() = %v, want %v", got, tc.want)
			}
			tearDown(t)
		})
	}
}
