package main

import "net/http"

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("GET /{$}", app.home)
	mux.HandleFunc("GET /experience", app.experience)
	mux.HandleFunc("GET /hobbies", app.hobbies)
	mux.HandleFunc("GET /contact", app.contact)

	return mux
}
