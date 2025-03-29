package main

import (
	"flag"
	"html/template"
	"log/slog"
	"net/http"
	"os"
	"sync"

	"github.com/go-playground/form"
)

type application struct {
	templateCache map[string]*template.Template
	logger        *slog.Logger
	formDecoder   *form.Decoder
	debug         bool
	mailer        *Mailer
	wg            sync.WaitGroup
}

func main() {
	//port := os.Getenv("PORT")
	port := flag.String("port", ":4000", "HTTP network port")
	debug := flag.Bool("debug", false, "enable debug mode in the ui")
	flag.Parse()
	email := os.Getenv("ARTGOMAIL")
	emailPass := os.Getenv("ARTGOMAILPASS")

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	cache, err := setTemplateCache()
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	formDecoder := form.NewDecoder()

	mailer, err := newMailer("smtp.gmail.com", 587, email, emailPass)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	app := &application{
		templateCache: cache,
		logger:        logger,
		formDecoder:   formDecoder,
		debug:         *debug,
		mailer:        mailer,
	}

	app.logger.Info("starting server", "port", port)
	err = http.ListenAndServe(*port, app.routes())
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
}
