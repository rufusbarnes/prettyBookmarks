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

	mux.Get("/", dynamicMiddleware.ThenFunc(app.home))
	mux.Get("/about", dynamicMiddleware.ThenFunc(app.about))

	// Bookmarks
	mux.Get("/bookmarks/:id", dynamicMiddleware.ThenFunc(app.getBookmarks))
	mux.Post("/bookmarks", dynamicMiddleware.ThenFunc(app.createBookmarks))
	mux.Put("/bookmarks/:id", dynamicMiddleware.ThenFunc(app.updateBookmarks))
	mux.Del("/bookmarks/:id", dynamicMiddleware.ThenFunc(app.deleteBookmarks))
	mux.Post("/bookmarks/import", dynamicMiddleware.ThenFunc(app.importBookmarks))

	// Static files
	fileServer := http.FileServer(http.Dir("./ui/static"))
	mux.Get("/static/", http.StripPrefix("/static", fileServer))

	return standardMiddleware.Then(mux)
}
