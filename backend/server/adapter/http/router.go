package http

import (
	"github.com/Mushus/trashbox/backend/server/adapter/http/handler"
	"github.com/Mushus/trashbox/backend/server/adapter/http/middleware"
	"github.com/Mushus/trashbox/backend/server/adapter/http/renderer"
	"github.com/Mushus/trashbox/backend/server/adapter/http/validator"
	"github.com/labstack/echo/v4"
	echomw "github.com/labstack/echo/v4/middleware"
)

func ProvideRouter(h handler.Handlers, v *validator.Validator, r *renderer.Renderer, s middleware.SessionMiddleware) (*echo.Echo, error) {
	e := echo.New()

	// settings
	e.HideBanner = true
	e.HidePort = true
	e.Validator = v
	e.Renderer = r

	logger := echomw.Logger()
	recover := echomw.Recover()
	e.Use(echo.MiddlewareFunc(s), logger, recover)

	// set up routings
	e.GET("/!/login", handler.Handlize(h.Session.GetLogin))
	e.POST("/!/login", handler.Handlize(h.Session.PostLogin))
	e.GET("/!/logout", handler.Handlize(h.Session.GetLogout))
	e.GET("/", handler.Handlize(h.Handler.GetIndex))
	e.GET("/:title", handler.Handlize(h.Handler.GetDocument))
	e.PUT("/:title", handler.Handlize(handler.Auth(h.Handler.PutDocument)))
	e.GET("/@/:id", handler.Handlize(h.Handler.GetAsset))
	e.GET("/@/:id/:format", handler.Handlize(h.Handler.GetFormatedAsset))
	e.POST("/@/", handler.Handlize(handler.Auth(h.Handler.UploadAsset)))

	return e, nil
}
