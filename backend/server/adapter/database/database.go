package database

import (
	"github.com/google/wire"

	"github.com/jinzhu/gorm"
	"golang.org/x/xerrors"
)

var DatabaseSet = wire.NewSet(
	ProvideDB,
	ProvideDocumentDatastore,
	ProvideUserDatastore,
)

// TxFunc トランザクション処理
type TxFunc func(tx *gorm.DB) error

// ProvideDB 指定した filepath をリポジトリにして、データベースを生成する
func ProvideDB() (*gorm.DB, func(), error) {
	filepath := "trashbox.db" // TODO: config
	db, err := gorm.Open("sqlite3", filepath)
	if err != nil {
		return nil, nil, xerrors.Errorf("failed to open database: %w", err)
	}

	cleanup := func() {
		db.Close()
	}

	return db, cleanup, nil
}

// Tx トランザクション処理 txFunc を実行する
func Tx(txFunc TxFunc, db *gorm.DB) (err error) {
	tx := db.Begin()
	// TODO: recover
	if tx.Error != nil {
		return tx.Error
	}
	defer func() {
		if r := recover(); r != nil {
			err = xerrors.Errorf("panic: %v", r)
		}
	}()

	if err := txFunc(tx); err != nil {
		// ロールバック処理は失敗してもリカバリできないので放置
		// ロールバック処理失敗しても時間で解決される
		tx.Rollback()
		return err
	}

	tx = tx.Commit()
	return nil
}
