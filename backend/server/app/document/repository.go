package document

import "github.com/Mushus/trashbox/backend/server/app/property"

type Repository interface {
	Get(id string) (*property.Document, error)
	Put(doc *property.Document) error
}
