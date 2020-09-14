package controllers

import "webvideos/060/views"

type Static struct {
	Home    *views.View
	Contact *views.View
}

func NewStatic() *Static {
	return &Static{
		Home:    views.NewView("bootstrap", "statics/home"),
		Contact: views.NewView("bootstrap", "statics/contact"),
	}
}
