package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/jonathanpetrone/aitarot/internal/server"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	serverAddr := ":8080"
	log.Printf("Starting server at port %s", serverAddr)

	mux := http.NewServeMux()

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

	err := httpServer.ListenAndServe()
	if err != nil {
		log.Printf("Server error: %v", err)
		log.Fatal(err)
	}
}
