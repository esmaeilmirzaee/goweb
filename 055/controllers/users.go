package controllers

import (
	"fmt"
	"net/http"
	"webvideos/055/views"

	"github.com/gorilla/schema"
)

type Users struct {
	NewView *views.View
}

type SignupForm struct {
	Email    string `schema:"email"`
	Password string `schema:"password"`
}

func (u *Users) Create(w http.ResponseWriter, r *http.Request) {
	var form SignupForm
	fmt.Fprintln(w, "User's created.")
	dec := schema.NewDecoder()
	if err := parseForm(r, &form); err != nil {
		panic(err)
	}
	fmt.Fprintln(w, form.Email)
}

func (u *Users) New(w http.ResponseWriter, r http.Request) {
	if err := u.NewView.Render(w, nil); err != nil {
		panic(err)
	}
}

func NewUsers() *Users {
	return &Users{
		NewView: view.NewView("bootstrap", "users/new")
	}
}
