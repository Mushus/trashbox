package database

import (
	"github.com/Mushus/trashbox/backend/server/app/property"
	"github.com/Mushus/trashbox/backend/server/app/user"
	"github.com/jinzhu/gorm"
)

type UserDatastore struct {
	db *gorm.DB
}

func ProvideUserDatastore(db *gorm.DB) user.Repository {
	return &UserDatastore{
		db: db,
	}
}

func (u UserDatastore) FindByLogin(login string) (*user.User, error) {
	var ud userData
	db := u.db.Where("`login` = ?", login).First(&ud)
	err := db.Error
	if gorm.IsRecordNotFoundError(err) {
		return nil, user.ErrUserNotFound
	} else if err != nil {
		return nil, err
	}

	return ud.toObject()
}

// Add 同一のログイン名の user がいない場合にユーザーを作成する
func (u UserDatastore) Add(added *user.User) error {
	ud := NewUserData(added)
	err := u.db.Create(&ud).Error
	if err != nil {
		return err
	}

	return nil
}

type userData struct {
	ID       int64  `db:"id"`
	Login    string `db:"login"`
	Password string `db:"password"`
}

func NewUserData(u *user.User) userData {
	return userData{
		ID:       u.ID(),
		Login:    u.Login(),
		Password: u.Password(),
	}
}

// TableName は userData の テーブル名を表す
func (userData) TableName() string {
	return "users"
}

func (u userData) toObject() (*user.User, error) {
	user, err := user.NewUser(property.User{
		ID:             u.ID,
		Login:          u.Login,
		HashedPassword: u.Password,
	})
	if err != nil {
		return nil, err
	}

	return &user, nil
}
