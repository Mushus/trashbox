package fs

import (
	"github.com/Mushus/trashbox/backend/server/app/property"
	"github.com/Mushus/trashbox/backend/server/app/repository"
)

type AssetDatastore struct {
}

func ProvideAssetDatastore() repository.Asset {
	return &AssetDatastore{}
}

func (a AssetDatastore) Get(id string) (*property.Asset, error) {
	return nil, nil
}
func (a AssetDatastore) Add(streamAsset property.Asset) (string, error) {
	return "", nil
}
func (a AssetDatastore) Remove(id string) error {
	return nil
}
