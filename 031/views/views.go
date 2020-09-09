package views

import "text/template"

func NewView(files ...string) *View {

	files = append(files, "views/footer.gohtml")
	t, err := template.ParseFiles(files...)
	if err != nil {
		panic(err)
	}

	return &View{
		Template: t,
	}
}

type View struct {
	Template *template.Template
}
