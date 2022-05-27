package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

func main() {

	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// Use the http.NewServeMux() function to initialize a new servemux (router), then
	// register the home function as the handler for the "/" URL pattern.
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	// Initialize a new http.Server struct.
	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  mux,
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
