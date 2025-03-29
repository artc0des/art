package main

import (
	"fmt"
	"net/http"

	"github.com/justinas/nosurf"
)

type emailForm struct {
	Email   string `form:"email"`
	Subject string `form:"subject"`
	Message string `form:"message"`
}

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, http.StatusOK, "home.html", "")

}
func (app *application) experience(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, http.StatusOK, "experience.html", "")
}
func (app *application) about(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, http.StatusOK, "about.html", "")
}
func (app *application) contact(w http.ResponseWriter, r *http.Request) {
	CSRFToken := nosurf.Token(r)
	app.render(w, r, http.StatusOK, "contact.html", CSRFToken)
}
func (app *application) contactForm(w http.ResponseWriter, r *http.Request) {
	var incomingForm emailForm

	err := app.decodePostForm(r, &incomingForm)
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	/*err = app.mailer.Send(incomingForm.Email, incomingForm.Subject, incomingForm.Message)
	if err != nil {
		app.serverError(w, r, err)
		return
	}*/
	app.background(func() {
		err = app.mailer.Send(incomingForm.Email, incomingForm.Subject, incomingForm.Message)
		if err != nil {
			app.serverError(w, r, err)
			return
		}
	})

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (app *application) render(w http.ResponseWriter, r *http.Request, status int, page string, data string) {
	ts, ok := app.templateCache[page]
	if !ok {
		err := fmt.Errorf("template does not exist in html folder")
		app.serverError(w, r, err)
		return
	}

	w.WriteHeader(status)
	err := ts.ExecuteTemplate(w, "index", data)
	if err != nil {
		app.serverError(w, r, err)
	}
}
