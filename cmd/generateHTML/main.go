package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	htmlhandler "github.com/jonathanpetrone/aitarot/internal/html-handler"
)

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

	fmt.Printf("🔮 Generating html template for %s %s %s...\n", sign, month, year)
	htmlhandler.MakeHTMLTemplate(sign, year, month)
}
