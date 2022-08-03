package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/dev-bimomure/go-pokemon/app/driver"
	"github.com/dev-bimomure/go-pokemon/app/models"
)

const version = "1.0.0"

type config struct {
	port   int
	env    string
	apiurl string
	db     struct {
		dsn string
	}
}

type application struct {
	config        config
	infolog       *log.Logger
	errorLog      *log.Logger
	templateCache map[string]*template.Template
	version       string
	DB            models.DBModel
}

func (app *application) serve() error {
	srv := &http.Server{
		Addr:              fmt.Sprintf("localhost:%d", app.config.port),
		Handler:           app.routes(),
		IdleTimeout:       30 * time.Second,
		ReadTimeout:       10 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
		WriteTimeout:      5 * time.Second,
	}

	app.infolog.Println(fmt.Sprintf("Starting HTTP server in %s mode on port %d", app.config.env, app.config.port))

	return srv.ListenAndServe()
}

func main() {
	var cfg config

	flag.IntVar(&cfg.port, "port", 8080, "Server port to listen on")
	flag.StringVar(&cfg.apiurl, "apiurl", "https://pokeapi.co/api/v2/", "Pokemon API url")
	flag.StringVar(&cfg.env, "env", "development", "Application enviorment {development|production}")
	flag.StringVar(&cfg.db.dsn, "dsn", "root:A19uxwkpfv96!@tcp(localhost:3306)/pokemon?parseTime=false", "DSN")
	flag.Parse()

	info_log := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	error_log := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	template_cache := make(map[string]*template.Template)

	conn, err := driver.OpenDB(cfg.db.dsn)
	if err != nil {
		error_log.Fatal(err)
	}
	defer conn.Close()

	app := &application{
		config:        cfg,
		infolog:       info_log,
		errorLog:      error_log,
		templateCache: template_cache,
		version:       version,
		DB:            models.DBModel{DB: conn},
	}

	err = app.serve()

	if err != nil {
		app.errorLog.Println(err)
		log.Fatal(err)
	}

}
