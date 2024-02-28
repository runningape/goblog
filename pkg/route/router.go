package route

import "github.com/gorilla/mux"

var Router *mux.Router

func Initialize() {
	Router = mux.NewRouter()
}

func Name2URL(routerName string, pair ...string) string {
	url, err := Router.Get(routerName).URL(pair...)
	if err != nil {
		//checkError(err)
		return ""
	}
	return url.String()
}
