// +build wireinject

package server

import (
	"github.com/Mushus/trashbox/backend/server/adapter/database"
	"github.com/Mushus/trashbox/backend/server/adapter/fs"
	"github.com/Mushus/trashbox/backend/server/adapter/http"
	"github.com/Mushus/trashbox/backend/server/app"
	"github.com/google/wire"
)

func InitializeServer() (*Server, func(), error) {
	wire.Build(
		database.DatabaseSet,
		fs.FSSet,
		http.HttpSet,
		app.AppSet,
		ProvideServer,
	)
	return nil, nil, nil
}
