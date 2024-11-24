package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
)

type application struct {
	templateCache map[string]*template.Template
}

func main() {
	port := os.Getenv("PORT")
	log.Printf("starting server on port %v", port)

	cache, err := setTemplateCache()
	if err != nil {
		log.Fatal(err.Error())
	}

	app := &application{
		templateCache: cache,
	}

	err = http.ListenAndServe(":"+port, app.routes())
	if err != nil {
		log.Fatal(err)
	}
}
