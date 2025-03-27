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
	dynamicMiddleware := alice.New(app.noSurf)
	mux.Handle("GET /contact", dynamicMiddleware.ThenFunc(app.contact))
	mux.Handle("POST /contact", dynamicMiddleware.ThenFunc(app.contactForm))

	standardChain := alice.New(app.recoverPanic, app.requestLogger, commonHeaders)

	return standardChain.Then(mux)
}
