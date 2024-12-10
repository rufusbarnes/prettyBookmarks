package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {
	// 1. Parse CLI Flags
	port := flag.String("port", "4040", "Specifies the port number on which the application should run")
	flag.Parse()

	// 2. Dependencies
	infoLog := log.New(os.Stdout, "[INFO]\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout, "[Error]\t", log.Ldate|log.Ltime)

	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
	}

	// 3. TLS Config

	// 4. Server Config

	mux := app.routes()

	// 5. Listen and Serve
	infoLog.Printf("Starting server on :%s", *port)
	err := http.ListenAndServe(("localhost:" + *port), mux)
	errorLog.Fatal(err)
}
