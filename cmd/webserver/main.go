package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/jonathanpetrone/aitarot/internal/auth"
	"github.com/jonathanpetrone/aitarot/internal/database"
	"github.com/jonathanpetrone/aitarot/internal/server"
	_ "github.com/lib/pq" // PostgreSQL driver
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

	// Load environment variables (already done in init, but keeping for clarity)
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

	// Initialize database queries
	db := database.New(dbConn)

	// Initialize auth services
	sessionService := auth.NewSessionService(db)
	authMiddleware := auth.NewAuthMiddleware(sessionService)

	// Start session cleanup goroutine
	go func() {
		ticker := time.NewTicker(24 * time.Hour) // Clean up daily
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				if err := sessionService.CleanupExpiredSessions(context.Background()); err != nil {
					log.Printf("Session cleanup failed: %v", err)
				}
			}
		}
	}()

	mux := http.NewServeMux()

	// Serve static assets
	mux.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("templates/assets"))))

	// Public routes
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		server.ServeStart(w, r, sessionService)
	})
	mux.HandleFunc("/home", server.ServeHome)
	mux.HandleFunc("/monthlyreadings", server.MonthlyReadingsHandler)
	mux.HandleFunc("/reading", server.ServeReading)
	mux.HandleFunc("/ask-1-card", server.ServeAskOneCard)
	mux.HandleFunc("/card-meaning", server.HandleCardMeaning)
	mux.HandleFunc("/askthetarot", server.ServeAskTheTarot)
	mux.HandleFunc("/login-user", server.ServeLoginUser)
	mux.HandleFunc("/register", server.ServeRegisterUser)

	// Authentication routes - UPDATED to pass sessionService
	mux.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		server.ServeAttemptLoginUser(w, r, db, sessionService)
	})
	mux.HandleFunc("/logout", func(w http.ResponseWriter, r *http.Request) {
		server.ServeLogout(w, r, sessionService)
	})
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
	mux.HandleFunc("/dashboard", authMiddleware.RequireAuth(server.ServeDashboard))
	mux.HandleFunc("/profile", authMiddleware.RequireAuth(server.ServeProfile))

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
