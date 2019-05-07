package fs

import (
	"github.com/Mushus/trashbox/backend/server/app/asset"
	"github.com/Mushus/trashbox/backend/server/app/property"
)

type AssetDatastore struct {
}

func ProvideAssetDatastore() asset.Repository {
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
