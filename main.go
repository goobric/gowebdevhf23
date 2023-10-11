package main

import (
	"html/template"
	"net/http"
)

var tpl *template.Template

func main() {
	// func ParseGlobal(pattern string)(*Template, error) Parse multi files at once
	// tpl, _ = template.ParseGlob("templates/*.html")
	// func (t *Template) ParseFiles(filenames ...string)(*Template, error)
	tpl, _ = tpl.ParseGlob("templates/*.html") // using the ParseFiles method on the Template Struct

	http.HandleFunc("/", indexHandler)      // register a Path to the indexHandler
	http.HandleFunc("/about", aboutHandler) // register a Path to the aboutHandler
	http.ListenAndServe(":8080", nil)       // start the server
}

func aboutHandlerHandler(w http.ResponseWriter, r *http.Request) {
	// func (t *Template) ExcuteTemplate(wr io.Writer, name string, data interface{}) error
	tpl.ExecuteTemplate(w, "about.html", nil)
}
