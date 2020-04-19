// +build ignore

package main

import (
	"errors"        // errors
	"html/template" //templating
	"io/ioutil"     // ioutilities
	"log"           // log stuff
	"net/http"      // perform http operations
	"regexp"        // regex all the things
)

// Page type
type Page struct {
	Title string
	Body  []byte
}

// The function template.Must is a convenience wrapper that panics when passed a non-nil error value, and otherwise returns the *Template unaltered. A panic is appropriate here; if the templates can't be loaded the only sensible thing to do is exit the program.
var templates = template.Must(template.ParseFiles("edit.html", "view.html"))

// MustCompile will panic if doesn't match, Compile won't panic
var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func getTitle(w http.ResponseWriter, r *http.Request) (string, error) {
	m := validPath.FindStringSubmatch(r.URL.Path)
	if m == nil {
		http.NotFound(w, r)
		return "", errors.New("invalid page title")
	}
	return m[2], nil // The title is second
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError) // 500
	}
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title, err := getTitle(w, r)
	if err != nil {
		return
	}
	p, err := loadPage(title) // loads title
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound) // this handles nil pages and redirects The http.Redirect function adds an HTTP status code of http.StatusFound (302) and a Location header to the HTTP response.
		return
	}
	renderTemplate(w, "view", p)
}

//editHandler loads the page (or, if it doesn't exist, create an empty Page struct), and displays an HTML form.
func editHandler(w http.ResponseWriter, r *http.Request) {
	title, err := getTitle(w, r)
	if err != nil {
		return
	}
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	renderTemplate(w, "edit", p)
}

// actually save our changes
func saveHandler(w http.ResponseWriter, r *http.Request) {
	title, err := getTitle(w, r)
	if err != nil {
		return
	}
	body := r.FormValue("body")
	p := &Page{Title: title, Body: []byte(body)} //  We must convert that value to []byte before it will fit into the Page struct. We use []byte(body) to perform the conversion.
	err = p.save()                               // save the changes
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound) // redirects on success
}

func main() {
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)
	http.HandleFunc("/save/", saveHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
