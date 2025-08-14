package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

type config struct {
	addr    string
	invoker string
}

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {
	var cfg config
	flag.StringVar(&cfg.addr, "addr", ":4001", "Port server will listen on")
	flag.StringVar(&cfg.invoker, "invoker", "Victory", "The person who started the application")
	flag.Parse()

	infoLog := log.New(os.Stdout, "[INFO]\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "[ERROR]\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
	}

	srv := &http.Server{
		Addr:     cfg.addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	infoLog.Printf("This application was started by %s", cfg.invoker)
	infoLog.Printf("Server will run on port %s.", cfg.addr)

	err := srv.ListenAndServe()

	infoLog.Printf("Server is running on port %s.", cfg.addr)
	errorLog.Fatal(err)
}
