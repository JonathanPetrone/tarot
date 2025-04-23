package server

import (
	"html/template"
	"net/http"
)

var Tmpl *template.Template
var DynamicContentTmpl *template.Template

func ServeStart(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	data := struct{ Title string }{Title: "AI Tarot"}
	tmpl.Execute(w, data)
}
