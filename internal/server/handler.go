package server

import (
	"fmt"
	"html/template"
	"net/http"
	"net/url"
	"regexp"
	"strings"

	"github.com/jonathanpetrone/aitarot/internal/tarot"
	"github.com/jonathanpetrone/aitarot/internal/timeutil"
)

func ServeStart(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	data := struct {
		Title string
		Year  string
		Month string
	}{
		Title: "AI Tarot",
		Year:  timeutil.CurrentTime.Year,
		Month: timeutil.CurrentTime.Month,
	}
	tmpl.Execute(w, data)
}

func ServeStartAdmin(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/admin_index.html"))
	data := struct{ Title string }{Title: "Admin page"}
	tmpl.Execute(w, data)
}

func ServeAdminCreateNewReadings(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/admin_new_readings.html"))
	tmpl.
		Execute(w, nil)
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

// zodiac
func MonthlyReadingsHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	year := query.Get("year")
	month := query.Get("month")

	tmpl := template.Must(template.ParseFiles("templates/monthly_readings.html"))

	data := struct {
		Year  string
		Month string
	}{
		Year:  year,
		Month: month,
	}

	tmpl.Execute(w, data)
}

/*
func MonthlyReadingsHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/zodiac_signs.html"))

	monthParam := r.URL.Query().Get("month") // e.g., "2025-05"
	parts := strings.Split(monthParam, "-")
	if len(parts) != 2 {
		http.Error(w, "Invalid month", http.StatusBadRequest)
		return
	}

	data := struct {
		Year  string
		Month string
	}{
		Year:  parts[0],
		Month: parts[1],
	}

	tmpl.Execute(w, data)
}
*/

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
			<p class="mb-4">{{.Career}}</p>
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

var cleanupPattern = regexp.MustCompile(`[^\p{L}\p{Zs}]+`) // keep only letters and spaces

func cleanCardName(raw string) string {
	// Decode URL encoding
	decoded, _ := url.QueryUnescape(raw)

	// Remove emoji, numbers, and position labels (e.g., "🌴 1. ")
	decoded = strings.TrimSpace(decoded)

	// Remove leading emoji and digits (e.g., "🌾 4. ")
	parts := strings.SplitN(decoded, "–", 2) // keep only before the dash
	cardPart := parts[0]

	// Remove anything that's not a letter or space
	clean := cleanupPattern.ReplaceAllString(cardPart, "")
	return strings.TrimSpace(clean)
}

func HandleCardMeaning(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	rawName := query.Get("name")
	id := query.Get("id")

	cardName := cleanCardName(rawName)

	meaning, ok := tarot.CardMeanings[cardName]
	if !ok {
		http.Error(w, fmt.Sprintf("Card not found: '%s'", cardName), http.StatusNotFound)
		return
	}

	data := struct {
		Meaning string
		Id      string
	}{
		Meaning: meaning.Description,
		Id:      id,
	}

	tmpl := template.Must(template.New("meaning").Parse(`
    <div id="general-meaning-{{ .Id }}" x-data="{ show: true }">
        <button 
            @click="show = !show"
            class="text-sm text-blue-400 underline hover:text-blue-600">
            <span x-text="show ? 'Hide General Card Meaning' : 'Show General Card Meaning'"></span>
        </button>
        <div x-show="show" x-transition class="text-pink-300 italic mt-2 mb-8">
            {{ .Meaning }}
        </div>
    </div>
`))

	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
