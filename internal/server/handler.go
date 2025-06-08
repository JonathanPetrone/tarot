package server

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"time"

	"github.com/jonathanpetrone/aitarot/internal/astrology"
	"github.com/jonathanpetrone/aitarot/internal/auth"
	"github.com/jonathanpetrone/aitarot/internal/database"
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
		Year          string
		Month         string
		YearCurrent   string
		MonthCurrent  string
		YearPast      string
		MonthPast     string
		YearUpcoming  string
		MonthUpcoming string
	}{
		Year:          year,
		Month:         month,
		YearCurrent:   timeutil.CurrentTime.Year,
		MonthCurrent:  timeutil.CurrentTime.Month,
		YearPast:      timeutil.Past.Year,
		MonthPast:     timeutil.Past.Month,
		YearUpcoming:  timeutil.Upcoming.Year,
		MonthUpcoming: timeutil.Upcoming.Month,
	}

	tmpl.Execute(w, data)
}

func ServeAskTheTarot(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/ask_the_tarot.html"))
	tmpl.Execute(w, nil)
}

func ServeLoginUser(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/login_user.html"))
	tmpl.Execute(w, nil)
}

func ServeRegisterUser(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/register.html"))
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

func ServeHealthCheck(w http.ResponseWriter, r *http.Request) {
	config, err := database.LoadConfigFromEnv()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := fmt.Sprintf(`Environment: %s
Host: %s
Port: %d
Database: %s
SSL Mode: %s`,
		config.Environment,
		config.Host,
		config.Port,
		config.Database,
		config.SSLMode)

	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(response))
}

func HandleRegisterUser(w http.ResponseWriter, r *http.Request, db *database.Queries) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse and validate
	email := strings.TrimSpace(strings.ToLower(r.FormValue("email")))
	password := r.FormValue("password")
	dateStr := r.FormValue("date_of_birth")

	if !isValidEmail(email) || !isValidPassword(password) {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Validate password using the auth package
	passwordErrors := auth.ValidatePassword(password, email)
	if len(passwordErrors) > 0 {
		// Join all error messages
		errorMsg := strings.Join(passwordErrors, "; ")
		http.Error(w, errorMsg, http.StatusBadRequest)
		return
	}

	// Hash password using the auth package
	hashedPassword, err := auth.HashPassword(password)
	if err != nil {
		log.Printf("Password hashing failed: %v", err)
		http.Error(w, "Registration failed", http.StatusInternalServerError)
		return
	}

	// Convert to sql.NullTime
	var birthDate sql.NullTime
	if dateStr != "" {
		parsed, err := time.Parse("2006-01-02", dateStr)
		if err != nil {
			http.Error(w, "Invalid date format", http.StatusBadRequest)
			return
		}
		birthDate = sql.NullTime{
			Time:  parsed,
			Valid: true,
		}
	} else {
		// If no date provided, set as null
		birthDate = sql.NullTime{Valid: false}
	}

	var zodiac database.NullZodiacSignEnum
	if birthDate.Valid {
		zodiacStr := astrology.GetZodiacSign(birthDate.Time)
		zodiac = database.NullZodiacSignEnum{
			ZodiacSignEnum: database.ZodiacSignEnum(zodiacStr),
			Valid:          true,
		}
	}

	_, err = db.CreateUser(r.Context(), database.CreateUserParams{
		Email:        email,
		PasswordHash: hashedPassword,
		DateOfBirth:  birthDate, // sql.NullTime
		Zodiac:       zodiac,    // You'll need to handle this too
	})

	if err != nil {
		log.Printf("User creation failed: %v", err)

		// Check for duplicate email more specifically
		if isDuplicateEmail(err) {
			log.Printf("Duplicate email registration attempt: %s", email)
			http.Error(w, "An account with this email already exists. Please try logging in instead.", http.StatusConflict)
			return
		}

		// Other database errors
		log.Printf("Unexpected database error during registration: %v", err)
		http.Error(w, "Registration failed. Please try again.", http.StatusInternalServerError)
		return
	}

	log.Printf("User registered successfully: %s with zodiac %s", email, zodiac.ZodiacSignEnum)
	http.Redirect(w, r, "/login-user", http.StatusSeeOther)
}

// Helper function to detect duplicate email errors
func isDuplicateEmail(err error) bool {
	errStr := strings.ToLower(err.Error())
	return strings.Contains(errStr, "duplicate") ||
		strings.Contains(errStr, "unique constraint") ||
		strings.Contains(errStr, "already exists") ||
		strings.Contains(errStr, "violates unique constraint")
}
