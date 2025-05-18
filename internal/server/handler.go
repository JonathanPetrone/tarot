package server

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"

	"github.com/jonathanpetrone/aitarot/internal/tarot"
)

func ServeStart(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	data := struct{ Title string }{Title: "AI Tarot"}
	tmpl.Execute(w, data)
}

func ServeStartAdmin(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/admin_index.html"))
	data := struct{ Title string }{Title: "Admin page"}
	tmpl.Execute(w, data)
}

func ServeAdminCreateNewReadings(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/admin_new_readings.html"))
	tmpl.Execute(w, nil)
}

func ServeAdminEditReadings(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/admin_edit_readings.html"))
	tmpl.Execute(w, nil)
}

func ServeAdminHome(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/admin_home.html"))
	tmpl.Execute(w, nil)
}

func ServeHome(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/home.html"))
	tmpl.Execute(w, nil)
}

func ZodiacGridHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/zodiac_signs.html"))
	tmpl.Execute(w, nil)
}

func ServeAskTheTarot(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/ask_the_tarot.html"))
	tmpl.Execute(w, nil)
}

func ServeAskOneCard(w http.ResponseWriter, r *http.Request) {
	card := tarot.DrawCards(1)[0] // Fetch the single card

	cardinfo := tarot.CardMeanings[card.Name]

	data := struct {
		CardName    string
		ImagePath   string
		Subheading  string
		Description string
		Love        string
		Career      string
	}{
		CardName:    card.Name,
		ImagePath:   card.ImagePath,
		Subheading:  cardinfo.Heading,
		Description: cardinfo.Description,
		Love:        cardinfo.Love,
		Career:      cardinfo.Career,
	}

	// Parse the inline template (only renders the card image)
	tmpl := template.Must(template.New("card").Parse(`
		<h2 class="text-white text-4xl text-center mb-6">{{.CardName}}</h2>
        <img src="{{.ImagePath}}" class="w-32 h-48 mb-6"/>
		<div class="flex flex-col">
			<h3 class="text-2xl mb-2">{{.Subheading}}</h3>
			<p class="mb-6">{{.Description}}</p>
			<h4 class="text-xl mb-2">Love:</h4>
			<p class="mb-6">{{.Love}}</p>
			<h4 class="text-xl mb-2">Career:</h4>
			<p>{{.Career}}</p>
		</div>
    `))

	// Render the template with the card data
	tmpl.Execute(w, data)
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
