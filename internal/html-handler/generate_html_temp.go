package htmlhandler

import (
	"fmt"
	"html/template"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

type MonthlyReading struct {
	Summary       string
	Cards         []Card
	ReadingStats  *Statistics
	FinalWhispers string
	Year          string
	Month         string
	Sign          string
}

type Card struct {
	Title         string
	Description   string
	Image         string
	Position      string
	SmallPosition string
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

var smallCardPositions = []string{
	"top-[40px] left-[24px]",
	"top-[40px] left-[24px] rotate-90",
	"top-[12px] left-[24px]",
	"top-[68px] left-[24px]",
	"top-[40px] left-[1px]",
	"top-[40px] left-[48px]",
	"top-[77px] left-[76px]",
	"top-[52px] left-[76px]",
	"top-[27px] left-[76px]",
	"top-[2px] left-[76px]",
}

func CapitalizeFirstCharacter(s string) string {
	if len(s) == 0 {
		return s // Return empty if string is empty
	}
	return strings.ToUpper(string(s[0])) + s[1:]
}

func getFilePathMonthlyReading(sign, year, month string) string {
	filePath := fmt.Sprintf("monthlyreadings/%s/%s/%s_2025.txt", year, month, sign)

	return filePath
}

func MakeHTMLTemplate(sign, year, month string) {
	chosenTemplate := "reading_template_02.html"
	filePathMR := getFilePathMonthlyReading(sign, year, month)

	stats, err := ParseStatistics(filePathMR)
	if err != nil {
		log.Fatal("Couldn't parse statistics")
	}
	cardsInReading, err := GetCardsFromReading(filePathMR)
	if err != nil {
		log.Fatal("Couldn't parse cards in reading")
	}

	/*
		fmt.Printf("Major Arcana Cards: %d\n", stats.MajorArcana)
		fmt.Printf("Minor Arcana Cards: %d\n", stats.MinorArcana)
		fmt.Printf("Most Common Suite: %s\n", strings.Join(stats.MostCommonSuit, ", "))
		fmt.Printf("Most Common Rank: %s\n", strings.Join(stats.MostCommonRank, ", "))
	*/

	// Given the params this picks up a Madame AI response
	filePathMadameAIresp := getFilePathMadameAI(sign, year, month)
	content := ExtractContentFromResponse(filePathMadameAIresp)

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
		Year:          year,
		Month:         CapitalizeFirstCharacter(month),
		Sign:          CapitalizeFirstCharacter(sign),
	}

	//remove symbols
	re := regexp.MustCompile(`[\x{1F300}-\x{1F5FF}\x{1F600}-\x{1F64F}\x{1F680}-\x{1F6FF}\x{1F700}-\x{1F77F}\x{1F780}-\x{1F7FF}\x{1F800}-\x{1F8FF}\x{1F900}-\x{1F9FF}\x{1FA00}-\x{1FA6F}\x{1FA70}-\x{1FAFF}\x{2600}-\x{26FF}\x{2700}-\x{27BF}]`)

	cardSections := parts[1 : len(parts)-1]
	for i, part := range cardSections {
		lines := strings.SplitN(part, "\n", 2)
		if len(lines) != 2 {
			log.Fatalf("Failed to split section into title + description:\n%s", part)
		}
		fmt.Println(cardsInReading[i].ImagePath)
		card := Card{
			Title:         strings.TrimSpace(lines[0]),
			Description:   re.ReplaceAllString(strings.TrimSpace(lines[1]), ""),
			Image:         cardsInReading[i].ImagePath,
			Position:      cardPositions[i],
			SmallPosition: smallCardPositions[i],
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
