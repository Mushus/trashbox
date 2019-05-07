package user_test

import (
	"testing"

	"github.com/Mushus/trashbox/backend/server/app/property"
	"github.com/Mushus/trashbox/backend/server/app/user"
	userMock "github.com/Mushus/trashbox/backend/testutil/user"
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
		r := userMock.NewMockRepository(ctrl)

		var expectErr error
		var expectUser *user.User
		if !c.validUser {
			expectErr = user.ErrUserNotFound
		} else {
			u, _ := user.NewUser(resultUser)
			expectUser = &u
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
