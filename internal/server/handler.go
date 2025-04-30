package server

import (
	"fmt"
	"html/template"
	"net/http"

	aihandler "github.com/jonathanpetrone/aitarot/internal/ai-handler"
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

func ServeExample(w http.ResponseWriter, r *http.Request) {
	reading, err := aihandler.ParseMonthlyReading("input/reading.txt")
	if err != nil {
		http.Error(w, "Failed to load reading: "+err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl := template.Must(
		template.New("dynamic_content.html").
			Funcs(template.FuncMap{
				"add": add,
			}).
			ParseFiles("templates/dynamic_content.html"),
	)

	err = tmpl.ExecuteTemplate(w, "dynamic_content.html", reading)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
