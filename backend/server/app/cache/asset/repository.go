package cache

import (
	"github.com/Mushus/trashbox/backend/server/app/asset"
)

type Repository interface {
	GetCache(id, format string) (asset.Asset, error)
	PutCache(asset asset.Asset, format string) error
	PurgeAll(id string) error
}
