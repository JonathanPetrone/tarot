package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/jonathanpetrone/aitarot/internal/server"
)

// rand.Seed() If I want predicability for testing

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

	// Serve static assets (must come first and be more specific)
	mux.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("templates/assets"))))

	// Dynamic routes
	mux.HandleFunc("/reading", server.ServeReading)
	mux.HandleFunc("/readings", server.ZodiacGridHandler)
	mux.HandleFunc("/", server.ServeStart) // Generic fallback
	mux.HandleFunc("/home", server.ServeHome)
	mux.HandleFunc("/admin", server.ServeStartAdmin)
	mux.HandleFunc("/admin/createreadings", server.ServeAdminCreateNewReadings)
	mux.HandleFunc("/admin/editreadings", server.ServeAdminEditReadings)
	mux.HandleFunc("/admin/home", server.ServeAdminHome)

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
