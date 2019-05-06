package database

import (
	"github.com/Mushus/trashbox/backend/server/app/property"
	"github.com/Mushus/trashbox/backend/server/app/repository"
	"github.com/jinzhu/gorm"
)

type UserDatastore struct {
	db *gorm.DB
}

func ProvideUserDatastore(db *gorm.DB) repository.User {
	return &UserDatastore{
		db: db,
	}
}

func (u UserDatastore) FindByLogin(login string) (*property.User, error) {
	var user userData
	db := u.db.Where("`login` = ?", login).First(&user)
	err := db.Error
	if gorm.IsRecordNotFoundError(err) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return user.toProp(), nil
}

// Add 同一のログイン名の user がいない場合にユーザーを作成する
func (u UserDatastore) Add(user *property.User) error {
	return nil
}

type userData struct {
	ID       int64  `db:"id"`
	Login    string `db:"login"`
	Password string `db:"password"`
}

func (u *userData) toProp() *property.User {
	if u == nil {
		return nil
	}
	return &property.User{
		ID:       u.ID,
		Login:    u.Login,
		Password: u.Password,
	}
}
