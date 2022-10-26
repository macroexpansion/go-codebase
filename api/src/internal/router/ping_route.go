package router

import (
	"github.com/gorilla/mux"

	c "api/src/internal/controller"
)

func pingRoute(router *mux.Router, c *c.Controller) {
	router.HandleFunc("/ping", c.Ping).Methods("GET")
}
