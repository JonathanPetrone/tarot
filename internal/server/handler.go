package server

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"

	aihandler "github.com/jonathanpetrone/aitarot/internal/ai-handler"
	"github.com/jonathanpetrone/aitarot/internal/astrology"
	"github.com/jonathanpetrone/aitarot/internal/tarot"
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

func ServeReading(w http.ResponseWriter, r *http.Request) {
	starSignString := r.URL.Query().Get("type")

	// Validate starsign first
	sign, ok := astrology.StarSignMap[starSignString]
	if !ok {
		http.Error(w, "Error: unknown star sign", http.StatusBadRequest)
		return
	}

	// Generate the spread
	spread := tarot.ReadSpread(tarot.CelticCross)

	// Format the reading
	reading := tarot.FormatReading(tarot.CelticCross, spread, sign, true)

	// Send the reading back
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(reading))
	fmt.Println(starSignString)
}

func ServeReadingJSON(w http.ResponseWriter, r *http.Request) {
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
