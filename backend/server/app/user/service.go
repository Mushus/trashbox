package user

import (
	"github.com/Mushus/trashbox/backend/server/app/property"
	"github.com/Mushus/trashbox/backend/server/app/repository"
	"golang.org/x/xerrors"
)

var (
	ErrInvalidLoginOrPassword = xerrors.New("invalid login or password")
)

// Service ユーザーサービス
type Service struct {
	repository repository.User
}

// ProvideService ユーザーサービスを作成する
func ProvideService(u repository.User) *Service {
	return &Service{
		repository: u,
	}
}

// VerifyUser ユーザーが正しいか検証する
// User が正しい場合はそのユーザーを返す
func (s Service) VerifyUser(login, password string) (*User, error) {
	userProp, err := s.repository.FindByLogin(login)
	if err != nil {
		if xerrors.Is(err, repository.ErrUserNotFound) {
			return nil, ErrInvalidLoginOrPassword
		}
		return nil, err
	}

	user, err := NewUser(*userProp)
	if err != nil {
		// NOTE: 実装上ここには到達しない
		return nil, nil
	}

	if !user.VerifyPassword(password) {
		// NOTE: サイドチャンネル攻撃に弱い気がする
		return nil, ErrInvalidLoginOrPassword
	}

	return user, nil
}

func (s Service) AddUser(user *property.User) error {
	return nil
}
