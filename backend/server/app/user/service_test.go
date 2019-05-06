package user_test

import (
	"testing"

	"github.com/Mushus/trashbox/backend/server/app/property"
	"github.com/Mushus/trashbox/backend/server/app/repository"
	"github.com/Mushus/trashbox/backend/server/app/user"
	"github.com/golang/mock/gomock"
	"golang.org/x/xerrors"
)

func TestServiceVerifyUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	resultUser := property.User{Login: "user", Password: "password"}

	cases := []struct {
		login         string
		password      string
		validUser     bool
		validPassword bool
	}{
		{
			login:         "user",
			password:      "password",
			validUser:     true,
			validPassword: true,
		},
		{
			login:         "user",
			password:      "invald_password",
			validUser:     true,
			validPassword: false,
		},
		{
			login:         "invalid_user",
			password:      "password",
			validUser:     false,
			validPassword: false,
		},
	}

	for _, c := range cases {
		r := repository.NewMockUser(ctrl)

		var expectErr error
		var expectUser property.User
		if !c.validUser {
			expectErr = repository.ErrUserNotFound
		} else {
			expectUser = resultUser
		}

		r.
			EXPECT().
			FindByLogin(c.login).
			Return(expectUser, expectErr)

		s := user.ProvideService(r)

		u, err := s.VerifyUser(c.login, c.password)
		if xerrors.Is(err, user.ErrInvalidLoginOrPassword) {

		} else if err != nil {
			t.Fatalf("verificate user error: %v", err)
		}

		actualVerify := u != nil
		expectVerify := c.validPassword && c.validUser
		if actualVerify != expectVerify {
			t.Fatalf("verificate user: expect: %v, actual: %v", expectVerify, actualVerify)
		}
	}
}
