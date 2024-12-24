package main

import (
	"html/template"
	"log/slog"
	"net/http"
	"os"
)

type application struct {
	templateCache map[string]*template.Template
	logger        *slog.Logger
}

func main() {
	//port := os.Getenv("PORT")
	port := "4000"

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	cache, err := setTemplateCache()
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	app := &application{
		templateCache: cache,
		logger:        logger,
	}

	app.logger.Info("starting server", "port", port)
	err = http.ListenAndServe(":"+port, app.routes())
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
}
