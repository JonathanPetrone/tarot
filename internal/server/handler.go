package server

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"
)

var Tmpl *template.Template
var ZodiacTmpl *template.Template
var DynamicContentTmpl *template.Template

func init() {
	// Initialize the dynamic content template on server startup
	var err error
	ZodiacTmpl, err = template.ParseFiles("templates/zodiac_signs.html")
	if err != nil {
		fmt.Println("Error parsing zodiac_signs.html:", err)
		// Handle the error appropriately, maybe panic in production
	}
}

func ServeStart(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	data := struct{ Title string }{Title: "AI Tarot"}
	tmpl.Execute(w, data)
}

func ServeHome(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/welcome.html"))
	tmpl.Execute(w, nil)
}

func ZodiacGridHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/zodiac_signs.html"))
	tmpl.Execute(w, nil)
}

func add(a, b int) int {
	return a + b
}

func ServeReading(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	sign := query.Get("sign")
	year := query.Get("year")
	month := query.Get("month")

	// Validate required query parameters
	if sign == "" || year == "" || month == "" {
		http.Error(w, "Missing query parameters: sign, year, and month are required.", http.StatusBadRequest)
		return
	}

	// Build the template path
	templateToParse := fmt.Sprintf("templates/readings/%s/%s/%s_%s_%s.html",
		year, month, strings.ToLower(sign), year, strings.ToLower(month))

	// Parse and execute the template
	tmpl, err := template.ParseFiles(templateToParse)
	if err != nil {
		http.Error(w, fmt.Sprintf("Template not found: %v", err), http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, fmt.Sprintf("Failed to render template: %v", err), http.StatusInternalServerError)
		return
	}
}
