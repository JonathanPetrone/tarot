package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
	aihandler "github.com/jonathanpetrone/aitarot/internal/ai-handler"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("❌ Failed to load .env file")
	}
}

func main() {
	if len(os.Args) < 4 {
		fmt.Printf("Usage: %s <sign> <year> <month>\n", os.Args[0])
		os.Exit(1)
	}

	sign := strings.ToLower(os.Args[1])
	year := os.Args[2]
	month := strings.ToLower(os.Args[3])

	if sign == "" || year == "" || month == "" {
		log.Fatal("❌ All of <sign>, <year>, and <month> must be provided")
	}

	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		log.Fatal("❌ OPENAI_API_KEY is not set")
	}

	filename := fmt.Sprintf("%s_%s.txt", sign, year)
	fmt.Printf("🔮 Generating reading for %s %s %s...\n", sign, month, year)
	aihandler.GetAIReading(apiKey, filename, year, month)
}
