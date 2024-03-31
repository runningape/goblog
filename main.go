package main

import (
	"net/http"

	"github.com/runningape/goblog/app/http/middlewares"
	"github.com/runningape/goblog/bootstrap"
	"github.com/runningape/goblog/config"
	c "github.com/runningape/goblog/pkg/config"
)

func init() {
	config.Initialize()
}

func main() {
	bootstrap.SetupDB()
	router := bootstrap.SetupRoute()

	http.ListenAndServe(":"+c.GetString("app.port"), middlewares.RemoveTrailingSlash(router))
}
