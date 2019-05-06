package http

import (
	"github.com/Mushus/trashbox/backend/server/adapter/http/handler"
	"github.com/Mushus/trashbox/backend/server/adapter/http/middleware"
	"github.com/Mushus/trashbox/backend/server/adapter/http/renderer"
	"github.com/Mushus/trashbox/backend/server/adapter/http/template"
	"github.com/Mushus/trashbox/backend/server/adapter/http/validator"
	"github.com/google/wire"
)

var HttpSet = wire.NewSet(
	middleware.NewSession,
	template.ProvideTemplates,
	renderer.ProvideRenderer,
	handler.ProvideSession,
	handler.ProvideHandler,
	validator.ProvideValidator,
	ProvideRouter,
	wire.Struct(new(handler.Handlers), "*"),
)
