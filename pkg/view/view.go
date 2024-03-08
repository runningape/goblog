package view

import (
	"html/template"
	"io"
	"path/filepath"
	"strings"

	"github.com/runningape/goblog/logger"
	"github.com/runningape/goblog/pkg/route"
)

func Render(w io.Writer, data interface{}, tplFiles ...string) {
	viewDir := "resources/views/"

	for i, f := range tplFiles {
		tplFiles[i] = viewDir + strings.Replace(f, ".", "/", -1) + ".html"
	}

	layoutsFiles, err := filepath.Glob(viewDir + "layouts/*.html")
	logger.LogError(err)

	allFiles := append(layoutsFiles, tplFiles...)

	tmpl, err := template.New("").
		Funcs(template.FuncMap{
			"RouteNameToURL": route.Name2URL,
		}).ParseFiles(allFiles...)
	logger.LogError(err)

	err = tmpl.ExecuteTemplate(w, "app", data)
	logger.LogError(err)

}
