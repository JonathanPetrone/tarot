package admin

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	aihandler "github.com/jonathanpetrone/aitarot/internal/ai-handler"
	"github.com/jonathanpetrone/aitarot/internal/astrology"
	htmlhandler "github.com/jonathanpetrone/aitarot/internal/html-handler"
	"github.com/jonathanpetrone/aitarot/internal/readings"
)

func CreateNewReadings(sign, year, month string) {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal("❌ Failed to get working directory")
	}

	envPath := fmt.Sprintf("%s/.env", wd)
	err = godotenv.Load(envPath)
	if err != nil {
		log.Fatalf("❌ Error loading .env file from %s", envPath)
	}

	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		log.Fatal("❌ OPENAI_API_KEY is not set")
	}

	if sign == "all" {
		// First part of creating the reading i.e drawing cards
		for _, sign := range astrology.ZodiacSigns {
			reading := readings.CreateReading(year, month, sign.Name)
			err := readings.FormatReadingForAI(reading)
			if err != nil {
				log.Printf("⚠️ Failed to save reading for %s: %v", sign.Name, err)
			}
		}

		// Second part of creating the reading i.e sending the drawn cards to MadameAI
		for _, sign := range astrology.ZodiacSigns {
			filenameDrawnCards := fmt.Sprintf("%s_%s.txt", sign.Name, year)
			fmt.Printf("🔮 Generating reading for %s %s %s...\n", sign, month, year)
			aihandler.GetAIReading(apiKey, filenameDrawnCards, year, month)
		}

		// Third part of creating the reading i.e making the html template
		for _, sign := range astrology.ZodiacSigns {
			fmt.Printf("🔮 Generating html template for %s %s %s...\n", sign, month, year)
			htmlhandler.MakeHTMLTemplate(sign.Name, year, month)
		}
	}

	// First part of creating the reading i.e drawing cards
	drawnCards := readings.CreateReading(year, month, sign)
	err = readings.FormatReadingForAI(drawnCards)
	if err != nil {
		log.Printf("⚠️ Failed to save reading for %s: %v", sign, err)
	}

	// Second part of creating the reading i.e sending the drawn cards to MadameAI
	filenameDrawnCards := fmt.Sprintf("%s_%s.txt", sign, year)
	fmt.Printf("🔮 Generating reading for %s %s %s...\n", sign, month, year)
	aihandler.GetAIReading(apiKey, filenameDrawnCards, year, month)

	// Third part of creating the reading i.e making the html template
	fmt.Printf("🔮 Generating html template for %s %s %s...\n", sign, month, year)
	htmlhandler.MakeHTMLTemplate(sign, year, month)
}
