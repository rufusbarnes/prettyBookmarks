package main

import (
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	standardMiddleware := alice.New(app.recoverPanic, app.logRequest, secureHeaders)
	dynamicMiddleware := alice.New(app.session.Enable, app.noSurf)

	mux := pat.New()

	// Bookmarks
	mux.Get("/", dynamicMiddleware.ThenFunc(app.home))

	// Static files
	fileServer := http.FileServer(http.Dir("./ui/static"))
	mux.Get("/static/", http.StripPrefix("/static", fileServer))

	mux.Get("/about", dynamicMiddleware.ThenFunc(app.about))

	return standardMiddleware.Then(mux)
}
