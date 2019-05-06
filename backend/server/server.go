package server

import (
	"github.com/labstack/echo/v4"
)

// Server サーバーインスタンスです
type Server struct {
	router *echo.Echo
}

// ProvideServer サーバーを作成する
func ProvideServer(router *echo.Echo) *Server {
	return &Server{
		router: router,
	}
}

// Start サーバーを起動します
func (s Server) Start() {
	addr := ":8080"

	r := s.router
	// let's start
	if err := r.Start(addr); err != nil {
		r.Logger.Fatal(err)
	}
}
