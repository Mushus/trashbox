package asset

import (
	"github.com/Mushus/trashbox/backend/server/app/property"
)

type Repository interface {
	Get(id string) (*property.Asset, error)
	Add(streamAsset property.Asset) (string, error)
	Remove(id string) error
}
