package route

import (
	"github.com/gorilla/mux"
	"github.com/runningape/goblog/routes"
)

var Router *mux.Router

func Initialize() {
	Router = mux.NewRouter()
	routes.RegisterWebRoutes(Router)
}

func Name2URL(routerName string, pair ...string) string {
	url, err := Router.Get(routerName).URL(pair...)
	if err != nil {
		//checkError(err)
		return ""
	}
	return url.String()
}
