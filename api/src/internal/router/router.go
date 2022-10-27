package router

import (
	"github.com/gorilla/mux"
	"gorm.io/gorm"

	c "api/src/internal/controller"
	"api/src/internal/middleware"

	"pgsql/repos"
)

func Router(db *gorm.DB) *mux.Router {
	router := mux.NewRouter()
	router.Use(middleware.ContentTypeApplicationJsonMiddleware)

	controller := c.NewController()
	pingRoute(router, controller)

	accountRepo := repos.NewAccountRepo(db)
	authController := c.NewAuthController(accountRepo)
	authRoute(router, authController)

	return router
}
