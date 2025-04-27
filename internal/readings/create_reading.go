package readings

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/jonathanpetrone/aitarot/internal/astrology"
	"github.com/jonathanpetrone/aitarot/internal/tarot"
)

type Reading struct {
	Year   int
	Month  string
	Zodiac astrology.StarSign
	Cards  []tarot.SpreadCard
}

func CreateReading(year int, month string, zodiac string) Reading {
	reading := Reading{
		Year:   year,
		Month:  month,
		Zodiac: astrology.StarSignMap[zodiac],
		Cards:  tarot.ReadSpread(tarot.CelticCross),
	}

	return reading
}

func FormatReadingForAI(r Reading) error {
	// Build the content
	var sb strings.Builder

	// Meta info
	metaReading := fmt.Sprintf("%v Monthly Reading for %s %d\n", r.Zodiac.Name, r.Month, r.Year)
	sb.WriteString(metaReading + "\n")

	// Cards drawn
	for _, position := range r.Cards {
		card := fmt.Sprintf("%2d. %-35s -> %s\n", position.Position, position.Context, position.Card.Name)
		sb.WriteString(card)
	}

	sb.WriteString("\n")

	// Analyze and add stats
	stats := tarot.Stats{}
	tarot.AnalyzeSpreadTarot(r.Cards, &stats)
	statsString := stats.String()
	sb.WriteString(statsString)

	// Final content as string
	content := sb.String()

	// Build file path
	dirPath := fmt.Sprintf("./monthlyreadings/%d/%s", r.Year, r.Month)
	fileName := fmt.Sprintf("%s_%d.txt", strings.ToLower(r.Zodiac.Name), r.Year)
	fullPath := filepath.Join(dirPath, fileName)

	// Create directory if it doesn't exist
	err := os.MkdirAll(dirPath, 0755)
	if err != nil {
		return fmt.Errorf("failed creating directory: %w", err)
	}

	// Write file
	err = os.WriteFile(fullPath, []byte(content), 0644)
	if err != nil {
		return fmt.Errorf("failed writing file: %w", err)
	}

	fmt.Println("Saved reading to", fullPath)
	return nil
}
