package main

import (
	"flag"
	"html/template"
	"log/slog"
	"net/http"
	"os"

	"github.com/go-playground/form"
)

type application struct {
	templateCache map[string]*template.Template
	logger        *slog.Logger
	formDecoder   *form.Decoder
	debug         bool
}

func main() {
	//port := os.Getenv("PORT")
	port := flag.String("port", ":4000", "HTTP network port")
	debug := flag.Bool("debug", false, "enable debug mode in the ui")
	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	cache, err := setTemplateCache()
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	formDecoder := form.NewDecoder()

	app := &application{
		templateCache: cache,
		logger:        logger,
		formDecoder:   formDecoder,
		debug:         *debug,
	}

	app.logger.Info("starting server", "port", port)
	err = http.ListenAndServe(*port, app.routes())
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
}
