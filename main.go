package main

import (
	"log"
	"net/http"

	"github.com/jonathanpetrone/aitarot/internal/astrology"
	"github.com/jonathanpetrone/aitarot/internal/server"
	"github.com/jonathanpetrone/aitarot/internal/tarot"
)

// rand.Seed() If I want predicability for testing

func tarotSpreads() {
	aries := astrology.StarSigns[0]
	spread := tarot.ReadSpread(tarot.CelticCross)
	spread2 := tarot.ReadSpread(tarot.PastPresentFuture)

	tarot.FormatReading(tarot.CelticCross, spread, aries, true)
	tarot.FormatReading(tarot.PastPresentFuture, spread2, aries, true)

	stats := tarot.Stats{}
	tarot.AnalyzeSpreadTarot(spread, &stats)
	stats.Print()
}

func main() {
	serverAddr := ":8081"
	log.Printf("Starting server at port %s", serverAddr)

	mux := http.NewServeMux()

	// Serve static assets (must come first and be more specific)
	mux.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("templates/assets"))))

	// Dynamic routes
	mux.HandleFunc("/reading", server.ServeReading)
	mux.HandleFunc("/readings", server.ZodiacGridHandler)
	mux.HandleFunc("/", server.ServeStart) // Generic fallback
	mux.HandleFunc("/home", server.ServeHome)
	mux.HandleFunc("/example", server.ServeExample)

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
