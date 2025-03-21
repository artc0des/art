package main

import (
	"net/http"

	"github.com/justinas/alice" // New import
)

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("GET /{$}", app.home)
	mux.HandleFunc("GET /experience", app.experience)
	mux.HandleFunc("GET /about", app.about)
	mux.HandleFunc("GET /contact", app.contact)

	standardChain := alice.New(commonHeaders, app.requestLogger)
	return standardChain.Then(mux)
}
