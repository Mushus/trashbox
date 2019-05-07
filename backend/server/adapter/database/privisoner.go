package database

import (
	"log"

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
		db := p.db.Exec(`CREATE TABLE IF NOT EXISTS users (
id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
login TEXT NOT NULL UNIQUE,
password TEXT NOT NULL
)`)
		if db.Error != nil {
			return xerrors.Errorf("failed to create user table: %w", db.Error)
		}
	}
	{
		defaultLogin := "user"
		defaultPassword := "user"
		if _, err := p.user.AddUserIfNotExists(defaultLogin, defaultPassword); err != nil {
			log.Println(err)
			return err
		}
	}
	return nil
}
