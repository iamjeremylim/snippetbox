package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

// Define a home handler function which writes a byte slice containing
// "Hello from Snippetbox" as the response body.
// The http.ResponseWriter parameter provides methods for assembling a HTTP
// response and sending it to the user.
// The *http.Request parameter is a pointer to a struct which holds information
// about the current request (like the HTTP method and the URL being requested)
// Check if the current request URL path exactly matches "/". If it doesn't, use
// the http.NotFound() function to send a 404 response to the client.
func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	files := []string{
		"./ui/html/home.page.tmpl",
		"./ui/html/base.layout.tmpl",
		"./ui/html/footer.partial.tmpl",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.errorLog.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	err = ts.Execute(w, nil)
	if err != nil {
		app.errorLog.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}

// Add a showSnippet handler function.
func (app *application) showSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
}

// Add a createSnippet handler function.
func (app *application) createSnippet(w http.ResponseWriter, r *http.Request) {
	// Use r.Method to check whether the request is using POST or not. Note that
	// http.MethodPost is a constant equal to the string "POST".
	if r.Method != http.MethodPost {
		// If it's not, use the Header().Set() method to add an 'Allow: POST' header
		// to the response header map. The first parameter is the header name, and
		// the second parameter is the header value.
		// Use the w.WriteHeader() method to send a 405 status code and the w.Write()
		// method to write a "Method Not Allowed" response body.
		w.Header().Set("Allow", http.MethodPost)
		// w.WriteHeader(405) +  w.Write([]byte("Method Not Allowed")) = line 41
		http.Error(w, "Method Not Allowed", 405)
		return
	}
	w.Write([]byte("Create a new snippet..."))
}
