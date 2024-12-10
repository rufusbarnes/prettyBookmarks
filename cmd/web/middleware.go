package main

import "net/http"

func (app *application) logRequest(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		app.infoLog.Printf("%v", r.URL.Path)
		next.ServeHTTP(w, r)
	}
}
