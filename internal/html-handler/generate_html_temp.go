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
	ReadingStats  *Statistics
	FinalWhispers string
}

type Card struct {
	Title       string
	Description string
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

func MakeHTMLTemplate(sign, year, month string) {
	chosenTemplate := "reading_template_02.html"
	filePath := fmt.Sprintf("monthlyreadings/%s/%s/%s_2025.txt", year, month, sign)

	stats, err := ParseStatistics(filePath)
	if err != nil {
		log.Fatal("Couldn't parse statistics")
	}
	cardsInReading, err := GetCardsFromReading(filePath)
	if err != nil {
		log.Fatal("Couldn't parse cards in reading")
	}

	fmt.Printf("Major Arcana Cards: %d\n", stats.MajorArcana)
	fmt.Printf("Minor Arcana Cards: %d\n", stats.MinorArcana)
	fmt.Printf("Most Common Suite: %s\n", strings.Join(stats.MostCommonSuit, ", "))
	fmt.Printf("Most Common Rank: %s\n", strings.Join(stats.MostCommonRank, ", "))

	// Given the params this picks up a Madame AI response
	content := ExtractContentFromResponse(sign, year, month)

	// Then we split it into parts
	parts, err := SplitMadameAIContent(content)
	if err != nil {
		log.Fatal(err)
	}

	// Then we store the content in a Monthly Reading struct
	reading := MonthlyReading{
		Summary:       strings.TrimSpace(parts[0]),
		FinalWhispers: strings.TrimSpace(parts[len(parts)-1]),
		ReadingStats:  &stats,
	}

	cardSections := parts[1 : len(parts)-1]
	for i, part := range cardSections {
		lines := strings.SplitN(part, "\n", 2)
		if len(lines) != 2 {
			log.Fatalf("Failed to split section into title + description:\n%s", part)
		}
		fmt.Println(cardsInReading[i].ImagePath)
		card := Card{
			Title:       strings.TrimSpace(lines[0]),
			Description: strings.TrimSpace(lines[1]),
			Image:       cardsInReading[i].ImagePath,
			Position:    cardPositions[i],
		}
		reading.Cards = append(reading.Cards, card)
	}

	// Load and render template
	// Resolve absolute template path
	rootDir, err := os.Getwd()
	if err != nil {
		log.Fatalf("❌ Failed to get working directory: %v", err)
	}
	tmplPath := filepath.Join(rootDir, "templates", chosenTemplate)
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
