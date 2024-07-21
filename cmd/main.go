package main

import (
	"net/http"

	"github.com/raafly/ewallet/internal/core"
	"github.com/raafly/ewallet/pkg/config"
	"github.com/raafly/ewallet/pkg/route"
)

func main() {
	db := config.NewConfig()
	usrRepo := core.NewUserRepo(db)
	usrService := core.NewUserService(usrRepo)
	usrHandler := core.NewUserHandler(usrService)
	route := route.NewRoute(usrHandler)
	// authMiddleware := middleware.NewAuthMiddleware(route)

	server := http.Server {
		Handler: route,
		Addr: "localhost:3000",
	}
	
	server.ListenAndServe()
}