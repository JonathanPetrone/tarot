package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/jonathanpetrone/aitarot/internal/astrology"
	"github.com/jonathanpetrone/aitarot/internal/readings"
)

func main() {
	// Expect: go run main.go 2025 May
	if len(os.Args) < 3 {
		log.Fatalf("Usage: %s <year> <month>", os.Args[0])
	}

	year := os.Args[1]
	month := strings.ToLower(os.Args[2])

	fmt.Printf("Generating readings for %s, %s...\n", month, year)

	for _, sign := range astrology.ZodiacSigns {
		reading := readings.CreateReading(year, month, sign.Name)
		readings.FormatReadingForAI(reading)
	}
}
