package controllers

import "webvideos/055/views"

type Static struct {
	Home    *views.View
	Contact *views.View
}

func NewStatric() *Static {
	return &Static{
		Home: views.NewView("bootstrap", "static/home"),
		Contact: views.NewView("bootstrap", "static/contact")
	}
}
