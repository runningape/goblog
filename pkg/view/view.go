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
	files, err := filepath.Glob(viewDir + "layouts/*.html")
	logger.LogError(err)

	name = strings.Replace(name, ".", "/", -1)
	newfiles := append(files, viewDir+name+".html")

	tmpl, err := template.New(name + ".html").
		Funcs(template.FuncMap{
			"RouteNameToURL": route.Name2URL,
		}).ParseFiles(newfiles...)
	logger.LogError(err)

	err = tmpl.ExecuteTemplate(w, "app", data)
	logger.LogError(err)
}
