package route

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/runningape/goblog/pkg/logger"
)

var route *mux.Router

func SetRoute(r *mux.Router) {
	route = r
}

func Name2URL(routerName string, pair ...string) string {
	url, err := route.Get(routerName).URL(pair...)
	if err != nil {
		logger.LogError(err)
		return ""
	}
	return url.String()
}

func GetRouteVariable(parameterName string, r *http.Request) string {
	vars := mux.Vars(r)
	return vars[parameterName]
}
