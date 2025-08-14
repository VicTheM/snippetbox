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

func main() {
	var cfg config
	flag.StringVar(&cfg.addr, "addr", ":4001", "Port server will listen on")
	flag.StringVar(&cfg.invoker, "invoker", "Victory", "The person who started the application")
	flag.Parse()

	infoLog := log.New(os.Stdout, "[INFO]\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "[ERROR]\t", log.Ldate|log.Ltime|log.Lshortfile)

	mux := http.NewServeMux()
	fileServer := http.FileServer(http.Dir("./ui/static/"))

	// Registering routes
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	infoLog.Printf("This application was started by %s", cfg.invoker)

	err := http.ListenAndServe(cfg.addr, mux)

	infoLog.Printf("Server is running on port %s.", cfg.addr)
	errorLog.Fatal(err)
}
