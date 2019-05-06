package handler

import (
	"net/http"

	"github.com/Mushus/trashbox/backend/server/adapter/http/middleware"

	"github.com/labstack/echo/v4"

	"github.com/Mushus/trashbox/backend/server/app/asset"
	"github.com/Mushus/trashbox/backend/server/app/document"
	"github.com/gorilla/sessions"
)

type Handlers struct {
	Session Session
	Handler Handler
}

// Handler handler
type Handler struct {
	document document.Repository
	asset    asset.Repository
}

// ProvideHandler ハンドラを生成する
func ProvideHandler(document document.Repository, asset asset.Repository) Handler {
	return Handler{
		document: document,
		asset:    asset,
	}
}

// Context represents the context of the current HTTP request
type Context struct {
	echo.Context
	UserID     int
	IsLoggedIn bool
}

// HandlerFunc is define a function to serve HTTP requests
type HandlerFunc func(Context) error

func Handlize(fn HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		userID := -1
		isLoggedIn := false

		sess, err := middleware.GetSession(c)
		if err == nil {
			userID = getUserIDFromSession(sess)
			isLoggedIn = userID != -1
		}

		return fn(Context{
			Context:    c,
			UserID:     userID,
			IsLoggedIn: isLoggedIn,
		})
	}
}

func Auth(next HandlerFunc) HandlerFunc {
	return func(c Context) error {
		if !c.IsLoggedIn {
			return c.String(http.StatusUnauthorized, "unauthorized")
		}
		if err := next(c); err != nil {
			return err
		}
		return nil
	}
}

func getUserIDFromSession(sess *sessions.Session) int {
	sessUesrID, ok := sess.Values[SessionKeyUserID]
	if !ok {
		return -1
	}
	userID, ok := sessUesrID.(int)
	if !ok || userID == 0 {
		return -1
	}
	return userID
}
