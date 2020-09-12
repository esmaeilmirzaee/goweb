package views

import (
	"net/http"
	"path/filepath"
	"text/template"
)

var (
	layoutDir string = "views/layouts/"
	layoutExt string = "gohtml"
)

func NewView(layout string, files ...string) *View {
	files = append(files, layoutFiles()...)
	t, err := template.ParseFiles(files...)
	if err != nil {
		panic(err)
	}
	return &View{
		Template: t,
		Layout:   layout,
	}
}

type View struct {
	Template *template.Template
	Layout   string
}

func (v *View) Render(w http.ResponseWriter, data interface{}) error {
	return v.Template.ExecuteTemplate(w, v.Layout, data)
}

func layoutFiles() []string {
	files, err := filepath.Glob(layoutDir + "*" + layoutExt)
	if err != nil {
		panic(err)
	}
	return files
}
