package views

import (
	"html/template"
	"path/filepath"
)

var (
	layoutDir string = "views/layout/"
	layoutExt string = ".gohtml"
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

func layoutFiles() []string {
	files, err := filepath.Glob(layoutDir + "*" + layoutExt)
	if err != nil {
		panic(err)
	}
	return files
}
