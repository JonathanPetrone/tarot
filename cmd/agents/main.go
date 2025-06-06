package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
	aihandler "github.com/jonathanpetrone/aitarot/internal/ai-handler"
	htmlhandler "github.com/jonathanpetrone/aitarot/internal/html-handler"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("❌ Failed to load .env file")
	}
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	if len(os.Args) < 4 {
		fmt.Printf("Usage: %s <sign> <year> <month>\n", os.Args[0])
		os.Exit(1)
	}

	agent := strings.ToLower(os.Args[1])
	sign := strings.ToLower(os.Args[2])
	year := os.Args[3]
	month := strings.ToLower(os.Args[4])

	if agent == "" || sign == "" || year == "" || month == "" {
		log.Fatal("❌ All of <agent>, <sign>, <year>, and <month> must be provided")
	}

	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		log.Fatal("❌ OPENAI_API_KEY is not set")
	}

	if agent == "madameai" {
		filename := fmt.Sprintf("%s_%s.txt", sign, year)
		fmt.Printf("🔮 Generating reading for %s %s %s...\n", sign, month, year)
		aihandler.GetAIReading(apiKey, filename, year, month)
	}

	if agent == "qualityagent" {
		fmt.Printf("🧹 Reviewing and updating text from... %s %s %s...\n", sign, month, year)
		aihandler.GetQualityReview(apiKey, year, month)

		err := htmlhandler.UpdateHTMLFromQualityAgent(year, month)
		if err != nil {
			log.Fatal("Failed to update HTML:", err)
		}

		log.Println("✅ Update process completed successfully")
	}
}
