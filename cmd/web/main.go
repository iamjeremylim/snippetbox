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

// Parsing runtime config settings for the app
// Establish dependencies for the handlers
// Running the HTTP server
func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
	}

	// Initialize a new http.Server struct.
	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	// The value returned from flag.String() function is a pointer to the flag
	// value, not the value itself. So we need to deference it before using it.
	// We are using Printf() to interpolate the address with the log msg.
	// If http.ListenAndServe() returns an error we use the log.Fatal() function
	// to log the error message and exit. Note that any error returned by
	// http.ListenAndServe() is always non-nil.
	infoLog.Printf("Starting server on %s", *addr)

	// Call the ListenAndServe() method on http.Server struct.
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}
