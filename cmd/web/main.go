package main

import (
	"html/template"
	"log"
	"net/http"
)

type application struct {
	templateCache map[string]*template.Template
}

func main() {
	log.Print("starting server on port 4000")

	cache, err := setTemplateCache()
	if err != nil {
		log.Fatal(err.Error())
	}

	app := &application{
		templateCache: cache,
	}

	err = http.ListenAndServe(":4000", app.routes())
	if err != nil {
		log.Fatal(err)
	}
}
