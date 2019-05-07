package user

import (
	"github.com/Mushus/trashbox/backend/server/app/property"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/xerrors"
)

// User ユーザー情報
type User struct {
	id    int64  `db:"id"`
	login string `db:"login"`
	// password express hashed password
	password string `db:"password"`
}

// NewUser ユーザーを作成する
func NewUser(prop property.User) (User, error) {
	u := User{
		id:       prop.ID,
		login:    prop.Login,
		password: prop.HashedPassword,
	}
	if prop.HashedPassword == "" {
		err := u.SetPassword(prop.Password)
		if err != nil {
			return u, xerrors.Errorf("failed to create user: %w", err)
		}
	}

	return u, nil
}

// SetPassword パスワードを password に変更する
func (u *User) SetPassword(password string) error {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return xerrors.Errorf("cannot generate user password: %w", err)
	}
	u.password = string(hashed)
	return nil
}

// VerifyPassword パスワードが正しいかどうか調べる
func (u User) VerifyPassword(password string) bool {
	comparePassword := []byte(password)
	hashedPassword := []byte(u.password)
	return bcrypt.CompareHashAndPassword(hashedPassword, comparePassword) == nil
}

// ToProp is convert usert to property.User
func (u *User) ToProp() *property.User {
	if u != nil {
		return nil
	}
	return &property.User{
		ID:             u.id,
		Login:          u.login,
		HashedPassword: u.password,
		Password:       "",
	}
}

// ID is users identifer
func (u *User) ID() int64 {
	return u.id
}

// Login is users login identier
func (u *User) Login() string {
	return u.login
}

// Password is users hased password
func (u *User) Password() string {
	return u.password
}
