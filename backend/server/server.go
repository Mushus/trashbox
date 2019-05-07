package server

import (
	"log"

	"github.com/Mushus/trashbox/backend/server/adapter/database"
	"github.com/labstack/echo/v4"
)

// Server サーバーインスタンスです
type Server struct {
	router      *echo.Echo
	provisioner *database.Provisioner
}

// ProvideServer サーバーを作成する
func ProvideServer(router *echo.Echo, provisioner *database.Provisioner) *Server {
	return &Server{
		router:      router,
		provisioner: provisioner,
	}
}

// Start サーバーを起動します
func (s Server) Start() {
	if err := s.provisioner.Privision(); err != nil {
		log.Printf("%v\n", err)
	}

	addr := ":8080"

	r := s.router
	// let's start
	if err := r.Start(addr); err != nil {
		r.Logger.Fatal(err)
	}
}
