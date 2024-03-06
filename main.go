package main

import (
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/runningape/goblog/app/http/middlewares"
	"github.com/runningape/goblog/bootstrap"
	"github.com/runningape/goblog/pkg/logger"
)

var router *mux.Router

func main() {
	bootstrap.SetupDB()
	router = bootstrap.SetupRoute()

	err := http.ListenAndServe(":3000", middlewares.RemoveTrailingSlash(router))
	logger.LogError(err)
}
