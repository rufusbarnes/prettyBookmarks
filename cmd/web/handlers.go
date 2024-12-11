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

func (app *application) getBookmarks(w http.ResponseWriter, r *http.Request) {}

func (app *application) createBookmarks(w http.ResponseWriter, r *http.Request) {}

func (app *application) updateBookmarks(w http.ResponseWriter, r *http.Request) {}

func (app *application) deleteBookmarks(w http.ResponseWriter, r *http.Request) {}

func (app *application) importBookmarks(w http.ResponseWriter, r *http.Request) {}
