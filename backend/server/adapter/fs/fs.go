package fs

import (
	"github.com/google/wire"
)

var FSSet = wire.NewSet(
	ProvideAssetDatastore,
)
