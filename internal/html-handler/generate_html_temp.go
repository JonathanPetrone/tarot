package htmlhandler

import (
	"fmt"
	"html/template"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type MonthlyReading struct {
	Summary       string
	Cards         []Card
	FinalWhispers string
}

type Card struct {
	Title       string
	Description string
	Icon        string
	Image       string
	Position    string
	SmallSpread []SmallCard
}

type SmallCard struct {
	Image    string
	Position string
}

var cardPositions = []string{
	"top-[164px] left-[110px]",
	"top-[164px] left-[110px] rotate-90",
	"top-[40px] left-[110px]",
	"top-[290px] left-[110px]",
	"top-[164px] left-[10px]",
	"top-[164px] left-[210px]",
	"top-[310px] left-[348px]",
	"top-[210px] left-[348px]",
	"top-[110px] left-[348px]",
	"top-[10px] left-[348px]",
}

// Parses and renders the HTML template
func MakeHTMLTemplate(sign, year, month string) {
	content := ExtractContentFromResponse(sign, year, month)

	parts, err := SplitMadameAIContent(content)
	if err != nil {
		log.Fatal(err)
	}

	reading := MonthlyReading{
		Summary:       strings.TrimSpace(parts[0]),
		FinalWhispers: strings.TrimSpace(parts[len(parts)-1]),
	}

	cardSections := parts[1 : len(parts)-1]
	for _, part := range cardSections {
		lines := strings.SplitN(part, "\n", 2)
		if len(lines) != 2 {
			log.Fatalf("Failed to split section into title + description:\n%s", part)
		}

		card := Card{
			Title:       strings.TrimSpace(lines[0]),
			Description: strings.TrimSpace(lines[1]),
		}
		reading.Cards = append(reading.Cards, card)
	}

	// Load and render template
	// Resolve absolute template path
	rootDir, err := os.Getwd()
	if err != nil {
		log.Fatalf("❌ Failed to get working directory: %v", err)
	}
	tmplPath := filepath.Join(rootDir, "templates", "reading_template.html")
	fmt.Println("📄 Attempting to load template from:", tmplPath)

	// Parse template with funcMap
	tmpl, err := template.New("reading").Funcs(template.FuncMap{
		"add": func(i, j int) int { return i + j },
	}).ParseFiles(tmplPath)
	if err != nil {
		log.Fatalf("❌ Failed to parse template: %v", err)
	}

	// Create output directory if needed
	outputDir := fmt.Sprintf("templates/readings/%s/%s", year, month)
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		log.Fatalf("Failed to create output folder: %v", err)
	}

	outputFile := fmt.Sprintf("%s_%s_%s.html", strings.ToLower(sign), year, strings.ToLower(month))
	outputPath := filepath.Join(outputDir, outputFile)

	f, err := os.Create(outputPath)
	if err != nil {
		log.Fatalf("Failed to create output file: %v", err)
	}
	defer f.Close()

	err = tmpl.ExecuteTemplate(f, "reading", reading)
	if err != nil {
		log.Fatalf("Failed to render HTML template: %v", err)
	}

	log.Printf("✅ Tarot reading HTML generated: %s\n", outputPath)
}
