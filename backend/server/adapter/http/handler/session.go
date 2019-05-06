package handler

import (
	"net/http"

	"github.com/Mushus/trashbox/backend/server/app"

	"github.com/Mushus/trashbox/backend/server/adapter/http/middleware"
	"github.com/Mushus/trashbox/backend/server/adapter/http/validator"

	"github.com/Mushus/trashbox/backend/server/adapter/http/template"
)

// SessionKeyUserID is used to obtain user ID from the session
const SessionKeyUserID = "userId"

// Session is the handler group for sessions
type Session struct {
	app *app.App
}

// ProvideSession provide the session handler
func ProvideSession(app *app.App) Session {
	return Session{
		app: app,
	}
}

// GetLogin ログインページ
func (s Session) GetLogin(c Context) error {
	return c.Render(http.StatusOK, template.TmplLogin, template.LoginView{
		Errors: validator.ValidationResult{},
	})
}

// LoginParam ログイン
type LoginParam struct {
	Login    string `form:"login" validate:"required"`
	Password string `form:"password" validate:"required"`
}

// PostLogin ログイン処理
func (s Session) PostLogin(c Context) error {
	var prm LoginParam
	if err := c.Bind(&prm); err != nil {
		return c.String(http.StatusBadRequest, "Bad Request")
	}

	if err := c.Validate(prm); err != nil {
		return c.Render(http.StatusOK, template.TmplLogin, template.LoginView{
			Errors: validator.ReportValidation(err),
		})
	}

	user, err := s.app.VerifyUser(prm.Login, prm.Password)
	if err != nil {
		return err
	}

	// success login
	sess, _ := middleware.GetSession(c)
	sess.Values[SessionKeyUserID] = user.ID()
	if err := middleware.SaveSession(c, sess); err != nil {
		return err
	}
	return c.Redirect(http.StatusSeeOther, "/")
}

// GetLogout is a handler to logout users
func (s Session) GetLogout(c Context) error {
	// TODO: logout process
	return c.Render(http.StatusOK, template.TmplLogout, nil)
}
