package user

import (
	"fmt"

	"github.com/Mushus/trashbox/backend/server/app/property"
	"golang.org/x/xerrors"
)

// Service ユーザーサービス
type Service struct {
	repository Repository
}

// ProvideService ユーザーサービスを作成する
func ProvideService(u Repository) *Service {
	return &Service{
		repository: u,
	}
}

// VerifyUser ユーザーが正しいか検証する
// User が正しい場合はそのユーザーを返す
func (s Service) VerifyUser(login, password string) (*User, error) {
	user, err := s.repository.FindByLogin(login)
	if err != nil {
		if xerrors.Is(err, ErrUserNotFound) {
			return nil, ErrInvalidLoginOrPassword
		}
		return nil, xerrors.Errorf("failed to find user: %w", err)
	}
	fmt.Printf("%#v", user)

	if !user.VerifyPassword(password) {
		// NOTE: サイドチャンネル攻撃に弱い気がする
		// 最大経過数待機か?
		return nil, ErrInvalidLoginOrPassword
	}

	return user, nil
}

// AddUser ユーザーを追加する
func (s Service) AddUser(user *User) error {
	return s.repository.Add(user)
}

func (s Service) AddUserIfNotExists(login, password string) (*User, error) {
	{
		// find exists user
		user, err := s.repository.FindByLogin(login)
		if user != nil {
			return user, nil
		}
		if !xerrors.Is(err, ErrUserNotFound) {
			return nil, xerrors.Errorf("failed to find user: %w", err)
		}
	}
	{
		// create new user
		user, err := NewUser(property.User{
			Login:    login,
			Password: password,
		})
		if err != nil {
			return nil, xerrors.Errorf("failed to create user: %w", err)
		}

		if err = s.repository.Add(&user); err != nil {
			return nil, xerrors.Errorf("failed to create user: %w", err)
		}

		return &user, nil
	}
}
