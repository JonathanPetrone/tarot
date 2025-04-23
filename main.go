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

	tarot.FormatReading(tarot.CelticCross, spread, aries)
	tarot.FormatReading(tarot.PastPresentFuture, spread2, aries)

	stats := tarot.Stats{}
	tarot.AnalyzeSpreadTarot(spread, &stats)
	stats.Print()
}

func main() {
	serverAddr := ":8080"
	log.Printf("Starting server at port %s", serverAddr)
	mux := http.NewServeMux()

	mux.Handle("GET /templates/", http.StripPrefix("/templates/", http.FileServer(http.Dir("templates"))))
	mux.Handle("GET /", http.HandlerFunc(server.ServeStart))

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
