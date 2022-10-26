package router

import (
	"github.com/gorilla/mux"

	c "api/src/internal/controller"
)

func authRoute(router *mux.Router, controller *c.AuthController) {
	router.HandleFunc("/auth/register", controller.Register).Methods("POST")
	router.HandleFunc("/auth/login", controller.Login).Methods("POST")
}
