package route

import "github.com/gorilla/mux"

var Router *mux.Router

func Initialize() {
	Router = mux.NewRouter()
}

func Name2URL(routerName string, pairs ...string) string {
	url, err := Router.Get(routerName).URL(pairs...)
	if err != nil {
		//checkError(err)
		return ""
	}
	return url.String()
}
