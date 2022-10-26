package router

import (
	"github.com/gorilla/mux"
	"gorm.io/gorm"

	c "api/src/internal/controller"
	"api/src/internal/middleware"
)

func Router(db *gorm.DB) *mux.Router {
	router := mux.NewRouter()
	router.Use(middleware.ContentTypeApplicationJsonMiddleware)

	controller := c.NewController()
	pingRoute(router, controller)

	authController := c.NewAuthController(db)
	authRoute(router, authController)

	return router
}
