package main

import (
	"log"

	"github.com/joho/godotenv"
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

	err := htmlhandler.UpdateHTMLFromQualityAgent("2025", "june")
	if err != nil {
		log.Fatal("Failed to update HTML:", err)
	}

	log.Println("✅ Update process completed successfully")
}
