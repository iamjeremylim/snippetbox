package main

import (
	"errors"
	"fmt"

	"net/http"
	"strconv"

	"iamjeremylim.net/snippetbox/pkg/models"
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
		app.notFound(w)
		return
	}

	s, err := app.snippets.Latest()
	if err != nil {
		app.serveError(w, err)
		return
	}

	app.render(w, r, "home.page.tmpl", &templateData{Snippets: s})

	// data := &templateData{Snippets: s}

	// files := []string{
	// 	"./ui/html/home.page.tmpl",
	// 	"./ui/html/base.layout.tmpl",
	// 	"./ui/html/footer.partial.tmpl",
	// }

	// ts, err := template.ParseFiles(files...)
	// if err != nil {
	// 	app.serveError(w, err)
	// 	return
	// }

	// err = ts.Execute(w, data)
	// if err != nil {
	// 	app.serveError(w, err)
	// }
}

// Add a showSnippet handler function.
func (app *application) showSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	s, err := app.snippets.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
		} else {
			app.serveError(w, err)
		}
		return
	}

	app.render(w, r, "show.page.tmpl", &templateData{Snippet: s})

	// data := &templateData{Snippet: s}

	// files := []string{
	// 	"./ui/html/show.page.tmpl",
	// 	"./ui/html/base.layout.tmpl",
	// 	"./ui/html/footer.partial.tmpl",
	// }

	// ts, err := template.ParseFiles(files...)
	// if err != nil {
	// 	app.serveError(w, err)
	// 	return
	// }

	// err = ts.Execute(w, data)
	// if err != nil {
	// 	app.serveError(w, err)
	// }
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
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	title := "0 snail"
	content := "O snail\nClimb Mount Fuji,\nBut slowly, slowly!\n\n– Kobayashi Issa"
	expires := "7"

	id, err := app.snippets.Insert(title, content, expires)
	if err != nil {
		app.serveError(w, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/snippet?id=%d", id), http.StatusSeeOther)
}
