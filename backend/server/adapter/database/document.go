package database

import (
	"github.com/Mushus/trashbox/backend/server/app/document"
	"github.com/Mushus/trashbox/backend/server/app/property"
	"github.com/jinzhu/gorm"
)

type DocumentDatastore struct {
	db *gorm.DB
}

func ProvideDocumentDatastore(db *gorm.DB) document.Repository {
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
