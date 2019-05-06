//go:generate mockgen -source ./repository.go -destination ./repository_mock.go -package repository

package repository

import (
	"github.com/Mushus/trashbox/backend/server/app/asset"
	"github.com/Mushus/trashbox/backend/server/app/property"
)

type User interface {
	FindByLogin(login string) (*property.User, error)
	Add(user *property.User) error
}

type Document interface {
	Get(id string) (*property.Document, error)
	Put(doc *property.Document) error
}

type Asset interface {
	Get(id string) (*property.Asset, error)
	Add(streamAsset property.Asset) (string, error)
	Remove(id string) error
}
type AssetCache interface {
	GetCache(id, format string) (asset.Asset, error)
	PutCache(asset asset.Asset, format string) error
	PurgeAll(id string) error
}
