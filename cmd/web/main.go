package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/VicTheM/snippetbox/internal/models"
	_ "github.com/go-sql-driver/mysql"
)

type config struct {
	addr    string
	invoker string
	dsn     string
}

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	snippets *models.SnippetModel
}

func main() {
	var cfg config
	flag.StringVar(&cfg.addr, "addr", ":4001", "Port server will listen on")
	flag.StringVar(&cfg.invoker, "invoker", "Victory", "The person who started the application")
	flag.StringVar(&cfg.dsn, "dsn", "web:@62453170Vic@/snippetbox?parseTime=true", "Database connection data source name")
	flag.Parse()

	infoLog := log.New(os.Stdout, "[INFO]\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "[ERROR]\t", log.Ldate|log.Ltime|log.Lshortfile)

	db, err := openDB(cfg.dsn)
	if err != nil {
		errorLog.Fatal(err)
	}

	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
		snippets: &models.SnippetModel{DB: db},
	}

	srv := &http.Server{
		Addr:     cfg.addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	infoLog.Printf("This application was started by %s", cfg.invoker)
	infoLog.Printf("Server will run on port %s.", cfg.addr)

	err = srv.ListenAndServe()

	infoLog.Printf("Server is running on port %s.", cfg.addr)
	errorLog.Fatal(err)
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err != nil {
			db.Close()
		}
	}()

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
