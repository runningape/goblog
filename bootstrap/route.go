package bootstrap

import (
	"github.com/gorilla/mux"
	"github.com/runningape/goblog/pkg/route"
	"github.com/runningape/goblog/routes"
)

func SetupRoute() *mux.Router {
	router := mux.NewRouter()
	routes.RegisterWebRoutes(router)
	route.SetRoute(router)
	return router
}
