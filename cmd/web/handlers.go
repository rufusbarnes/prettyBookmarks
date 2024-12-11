package main

import (
	"net/http"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "home.page.tmpl", &templateData{})
}

func (app *application) about(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "about.page.tmpl", nil)
}

// func (app *application) XXXXX(w http.ResponseWriter, r *http.Request) {}
// open bookmark data
// view json bookmark data
// import json bookmark data
// user - signup/form, login/form, logout
