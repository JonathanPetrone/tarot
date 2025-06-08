package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/jonathanpetrone/aitarot/internal/database"
	"github.com/jonathanpetrone/aitarot/internal/server"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func buildDatabaseURL() string {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	sslmode := os.Getenv("DB_SSLMODE")

	// Handle empty password
	if password == "" {
		return fmt.Sprintf("postgres://%s@%s:%s/%s?sslmode=%s",
			user, host, port, dbname, sslmode)
	}

	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		user, password, host, port, dbname, sslmode)
}

func main() {
	serverAddr := ":8080"
	log.Printf("Starting server at port %s", serverAddr)

	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Build database connection string
	dbURL := buildDatabaseURL()

	dbConn, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer dbConn.Close()

	// Test the connection
	if err := dbConn.Ping(); err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}

	log.Println("Database connected successfully")

	mux := http.NewServeMux()
	db := database.New(dbConn)

	// Serve static assets
	mux.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("templates/assets"))))

	// Public routes
	mux.HandleFunc("/", server.ServeStart)
	mux.HandleFunc("/home", server.ServeHome)
	mux.HandleFunc("/monthlyreadings", server.MonthlyReadingsHandler)
	mux.HandleFunc("/reading", server.ServeReading)
	mux.HandleFunc("/ask-1-card", server.ServeAskOneCard)
	mux.HandleFunc("/card-meaning", server.HandleCardMeaning)
	mux.HandleFunc("/askthetarot", server.ServeAskTheTarot)
	mux.HandleFunc("/login-user", server.ServeLoginUser)
	mux.HandleFunc("/register", server.ServeRegisterUser)
	mux.HandleFunc("/create-user", func(w http.ResponseWriter, r *http.Request) {
		server.HandleRegisterUser(w, r, db)
	})

	// Admin routes
	mux.HandleFunc("/admin", server.ServeStartAdmin)
	mux.HandleFunc("/admin/createreadings", server.ServeAdminCreateNewReadings)
	mux.HandleFunc("/admin/editreadings", server.ServeAdminEditReadings)
	mux.HandleFunc("/admin/home", server.ServeAdminHome)

	// Healthcheck
	mux.HandleFunc("/health", server.ServeHealthCheck)

	// Protected routes
	// mux.HandleFunc("/dashboard", authService.RequireAuth(server.ServeDashboard))

	httpServer := &http.Server{
		Handler: mux,
		Addr:    serverAddr,
	}

	err = httpServer.ListenAndServe()
	if err != nil {
		log.Printf("Server error: %v", err)
		log.Fatal(err)
	}
}
