package view

import (
	"io"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/runningape/goblog/pkg/logger"
	"github.com/runningape/goblog/pkg/route"
)

func Render(w io.Writer, name string, data interface{}) {
	viewDir := "resources/views/"
	name = strings.Replace(name, ".", "/", -1)

	files, err := filepath.Glob(viewDir + "layouts/*.html")
	logger.LogError(err)

	newFiles := append(files, viewDir+name+".html")

	tmpl, err := template.New(name + ".html").
		Funcs(template.FuncMap{
			"RouteNameToURL": route.Name2URL,
		}).ParseFiles(newFiles...)
	logger.LogError(err)

	err = tmpl.ExecuteTemplate(w, "app", data)
	logger.LogError(err)
}
