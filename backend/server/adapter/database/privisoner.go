package database

import (
	"github.com/Mushus/trashbox/backend/server/app/property"
	"github.com/Mushus/trashbox/backend/server/app/user"
	"github.com/jinzhu/gorm"
	"golang.org/x/xerrors"
)

type Provisioner struct {
	db   *gorm.DB
	user *user.Service
}

func ProvideProvisioner(db *gorm.DB, user *user.Service) *Provisioner {
	return &Provisioner{
		db:   db,
		user: user,
	}
}

// Privision 指定した filepath をリポジトリにして、データベースを生成する
func (p Provisioner) Privision() error {
	{
		err := p.db.Exec(`CREATE TABLE IF NOT EXISTS users (
id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
login TEXT NOT NULL UNIQUE,
password TEXT NOT NULL
)`)
		if err != nil {
			return xerrors.Errorf("failed to create user table: %w", err)
		}
	}
	{
		// デフォルトユーザーとして admin / admin を作成する
		u, err := user.NewUser(property.User{Login: "admin", Password: "admin"})
		if err != nil {
			return err
		}

		if err := p.user.AddUser(u.ToProp()); err != nil {
			return err
		}
	}
	return nil
}
