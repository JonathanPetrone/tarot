package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"
	aihandler "github.com/jonathanpetrone/aitarot/internal/ai-handler"
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
	serverAddr := ":8081"
	log.Printf("Starting server at port %s", serverAddr)

	_, err := aihandler.ParseMonthlyReading("/Users/jonathanpetrone/Github/AITarot/input/reading.txt")
	if err != nil {
		fmt.Println("Error parsing reading:", err)
		return
	}

	mux := http.NewServeMux()

	// Serve static assets (must come first and be more specific)
	mux.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("templates/assets"))))

	// Dynamic routes
	mux.HandleFunc("/reading", server.ServeExample)
	mux.HandleFunc("/readings", server.ZodiacGridHandler)
	mux.HandleFunc("/", server.ServeStart) // Generic fallback
	mux.HandleFunc("/home", server.ServeHome)

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
