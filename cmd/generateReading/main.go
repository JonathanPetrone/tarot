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
	if len(os.Args) != 3 {
		fmt.Printf("Usage: %s <year> <month>\n", os.Args[0])
		os.Exit(1)
	}

	year := os.Args[1]
	month := strings.ToLower(os.Args[2])

	if year == "" || month == "" {
		log.Fatal("❌ Both <year> and <month> must be provided.")
	}

	fmt.Printf("🔮 Generating readings for %s %s...\n", month, year)

	for _, sign := range astrology.ZodiacSigns {
		reading := readings.CreateReading(year, month, sign.Name)
		err := readings.FormatReadingForAI(reading)
		if err != nil {
			log.Printf("⚠️ Failed to save reading for %s: %v", sign.Name, err)
		}
	}
}
