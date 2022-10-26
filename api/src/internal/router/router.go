package router

import (
	"github.com/gorilla/mux"

	c "api/src/internal/controller"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	// controller := c.NewController()

	router.HandleFunc("/ping", c.Ping)

	return router
}
