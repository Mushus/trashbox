package app

import (
	"github.com/Mushus/trashbox/backend/server/app/user"
	"github.com/google/wire"
)

var AppSet = wire.NewSet(
	user.ProvideService,
	ProvideApp,
)

type App struct {
	userService *user.Service
}

func ProvideApp(userService *user.Service) *App {
	return &App{
		userService: userService,
	}
}

func (a App) VerifyUser(login, password string) (*user.User, error) {
	// TODO: バリデーション
	return a.userService.VerifyUser(login, password)
}
