//go:generate mockgen -source ./repository.go -destination ./repository_mock.go -package repository

package repository

import (
	"github.com/Mushus/trashbox/backend/server/app/asset"
)

type AssetCache interface {
	GetCache(id, format string) (asset.Asset, error)
	PutCache(asset asset.Asset, format string) error
	PurgeAll(id string) error
}
