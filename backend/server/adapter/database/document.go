package database

import (
	"github.com/Mushus/trashbox/backend/server/app/property"
	"github.com/Mushus/trashbox/backend/server/app/repository"
	"github.com/jinzhu/gorm"
)

type DocumentDatastore struct {
	db *gorm.DB
}

func ProvideDocumentDatastore(db *gorm.DB) repository.Document {
	return &DocumentDatastore{
		db: db,
	}
}

func (d DocumentDatastore) Get(id string) (*property.Document, error) {
	return nil, nil
}
func (d DocumentDatastore) Put(doc *property.Document) error {
	return nil
}
