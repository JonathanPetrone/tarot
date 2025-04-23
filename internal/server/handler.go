package server

import (
	"fmt"
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

func ServeReading(w http.ResponseWriter, r *http.Request) {
	spreadType := r.URL.Query().Get("type") // Get the value of the 'type' query parameter

	switch spreadType {
	case "celticcross":
		// In a real application, you'd fetch the Celtic Cross reading data
		readingData := struct{ Result string }{Result: "Displaying the Celtic Cross reading..."}
		// Execute the dynamic_content.html template with the reading data
		err := DynamicContentTmpl.Execute(w, readingData)
		if err != nil {
			http.Error(w, "Error executing dynamic_content.html template", http.StatusInternalServerError)
			fmt.Println("Error executing dynamic_content.html template:", err)
			return
		}
	case "threecard":
		// In a real application, you'd fetch the Three Card reading data
		readingData := struct{ Result string }{Result: "Displaying the Three Card reading..."}
		// Execute the dynamic_content.html template with the reading data
		err := DynamicContentTmpl.Execute(w, readingData)
		if err != nil {
			http.Error(w, "Error executing dynamic_content.html template", http.StatusInternalServerError)
			fmt.Println("Error executing dynamic_content.html template:", err)
			return
		}
	default:
		http.Error(w, "Invalid reading type specified", http.StatusBadRequest)
	}
}
