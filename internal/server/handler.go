package server

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"

	"github.com/jonathanpetrone/aitarot/internal/tarot"
)

var Tmpl *template.Template
var DynamicContentTmpl *template.Template

func init() {
	// Initialize the dynamic content template on server startup
	var err error
	DynamicContentTmpl, err = template.ParseFiles("templates/dynamic_content.html")
	if err != nil {
		fmt.Println("Error parsing dynamic_content.html:", err)
		// Handle the error appropriately, maybe panic in production
	}
}

func ServeStart(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	data := struct{ Title string }{Title: "AI Tarot"}
	tmpl.Execute(w, data)
}

func ZodiacGridHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/zodiac-grid.html"))
	tmpl.Execute(w, nil)
}

func ServeReading(w http.ResponseWriter, r *http.Request) {
	spreadType := r.URL.Query().Get("type")

	var spread []tarot.SpreadCard
	var err error

	switch spreadType {
	case "celticcross":
		spread = tarot.ReadSpread(tarot.CelticCross)
	case "threecard":
		spread = tarot.ReadSpread(tarot.PastPresentFuture)
	default:
		http.Error(w, "Invalid reading type specified", http.StatusBadRequest)
		return
	}

	// Marshal the spread data into JSON
	jsonData, err := json.MarshalIndent(spread, "", "  ") // Use MarshalIndent for pretty printing
	if err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
		return
	}

	// Set the Content-Type header to application/json
	w.Header().Set("Content-Type", "application/json")

	// Write the JSON data to the ResponseWriter
	w.Write(jsonData)
}
