package route

import (
	"net/http"

	"github.com/gorilla/mux"
)

var Router *mux.Router

func Name2URL(routerName string, pair ...string) string {
	url, err := Router.Get(routerName).URL(pair...)
	if err != nil {
		//checkError(err)
		return ""
	}
	return url.String()
}

func GetRouteVariable(parameterName string, r *http.Request) string {
	vars := mux.Vars(r)
	return vars[parameterName]
}
