package controllers

import (
	"fmt"
	"net/http"
	"webvideos/060/views"
)

type Users struct {
	NewView *views.View
}

type SignupForm struct {
	Name     string `schema:"name"`
	Email    string `schema:"email"`
	Password string `schema:"password"`
}

func NewUsers() *Users {
	return &Users{
		NewView: views.NewView("bootstrap", "users/new"),
	}
}

func (u *Users) New(w http.ResponseWriter, r *http.Request) {
	if err := u.NewView.Render(w, nil); err != nil {
		panic(err)
	}
}

func (u *Users) Create(w http.ResponseWriter, r *http.Request) {
	var form SignupForm
	fmt.Fprintln(w, "user created.")
	if err := parseForm(r, &form); err != nil {
		panic(err)
	}
	fmt.Fprint(w, form)
}
