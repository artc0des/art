package main

import (
	"fmt"
	"log"
	"net/http"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	app.render(w, http.StatusOK, "about.html")

}
func (app *application) experience(w http.ResponseWriter, r *http.Request) {
	app.render(w, http.StatusOK, "experience.html")
}
func (app *application) hobbies(w http.ResponseWriter, r *http.Request) {
	app.render(w, http.StatusOK, "hobbies.html")
}
func (app *application) contact(w http.ResponseWriter, r *http.Request) {
	app.render(w, http.StatusOK, "contact.html")
}

func (app *application) render(w http.ResponseWriter, status int, page string) {
	ts, ok := app.templateCache[page]
	if !ok {
		err := fmt.Errorf("template does not exist in html folder")
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(status)
	err := ts.ExecuteTemplate(w, "index", nil)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}
