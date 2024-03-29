package view

import (
	"html/template"
	"io"
	"path/filepath"
	"strings"

	"github.com/runningape/goblog/logger"
	"github.com/runningape/goblog/pkg/auth"
	"github.com/runningape/goblog/pkg/flash"
	"github.com/runningape/goblog/pkg/route"
)

type D map[string]interface{}

func Render(w io.Writer, data D, tplFiles ...string) {
	RenderTemplate(w, "app", data, tplFiles...)
}

func RenderSimple(w io.Writer, data D, tplFiles ...string) {
	RenderTemplate(w, "simple", data, tplFiles...)
}

func RenderTemplate(w io.Writer, name string, data D, tplFiles ...string) {
	data["isLogined"] = auth.Check()
	data["loginUser"] = auth.User
	data["flash"] = flash.All()

	allFiles := getTemplateFiles(tplFiles...)

	tmpl, err := template.New("").
		Funcs(template.FuncMap{
			"RouteNameToURL": route.Name2URL,
		}).ParseFiles(allFiles...)
	logger.LogError(err)

	err = tmpl.ExecuteTemplate(w, name, data)
	logger.LogError(err)

}

func getTemplateFiles(tplFiles ...string) []string {
	viewDir := "resources/views/"

	for i, f := range tplFiles {
		tplFiles[i] = viewDir + strings.Replace(f, ".", "/", -1) + ".html"
	}

	layoutFiles, err := filepath.Glob(viewDir + "layouts/*.html")
	logger.LogError(err)

	return append(layoutFiles, tplFiles...)
}
