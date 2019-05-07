package middleware

import (
	"github.com/averagesecurityguy/random"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/michaeljs1990/sqlitestore"
	"golang.org/x/xerrors"
)

// SessionMaxAge is max age of sessions
const SessionMaxAge = 60 * 60 * 24 * 100

type SessionMiddleware echo.MiddlewareFunc

// NewSession is coreate in-memory session
func NewSession() (SessionMiddleware, error) {
	secretKey, err := random.Token()
	if err != nil {
		return nil, xerrors.Errorf("cannot create session secret key: %w", err)
	}
	// TODO: this store has not feature of autoremove session
	store, err := sqlitestore.NewSqliteStore(":memory:", "sessions", "/", SessionMaxAge, []byte(secretKey))
	if err != nil {
		return nil, err
	}
	return SessionMiddleware(session.Middleware(store)), nil
}

// SessionModel represent a session instance
type SessionModel struct {
	UserID string
}

func GetSession(c echo.Context) (*sessions.Session, error) {
	return session.Get("session", c)
}

func SaveSession(c echo.Context, sess *sessions.Session) error {
	sess.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   SessionMaxAge,
		HttpOnly: true,
	}
	return sess.Save(c.Request(), c.Response())
}
