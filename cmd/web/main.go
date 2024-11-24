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
	//port := os.Getenv("PORT")
	port := "4000"
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
