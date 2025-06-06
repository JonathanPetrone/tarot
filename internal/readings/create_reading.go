package readings

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/jonathanpetrone/aitarot/internal/astrology"
	"github.com/jonathanpetrone/aitarot/internal/tarot"
)

type Reading struct {
	Year   string
	Month  string
	Zodiac astrology.Zodiac
	Cards  []tarot.SpreadCard
}

func CreateReading(year string, month string, zodiac string) Reading {
	return Reading{
		Year:   year,
		Month:  month,
		Zodiac: astrology.ZodiacSignMap[zodiac],
		Cards:  tarot.ReadSpread(tarot.CelticCross),
	}
}

func FormatReadingForAI(r Reading) error {
	// Validate required inputs
	if r.Zodiac.Name == "" {
		return errors.New("must provide a zodiac sign")
	}
	if r.Month == "" {
		return errors.New("must provide a month for the reading")
	}
	if r.Year == "" {
		return errors.New("must provide a year for the reading")
	}

	// Build reading text
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("%s Monthly Reading for %s %s\n\n", r.Zodiac.Name, r.Month, r.Year))

	for _, position := range r.Cards {
		sb.WriteString(fmt.Sprintf("%2d. %-35s -> %s\n", position.Position, position.Context, position.Card.Name))
	}
	sb.WriteString("\n")

	stats := tarot.Stats{}
	tarot.AnalyzeSpreadTarot(r.Cards, &stats)
	sb.WriteString(stats.String())

	// Final content
	content := sb.String()

	// Construct file path
	dirPath := filepath.Join("./monthlyreadings", r.Year, strings.ToLower(r.Month))
	fileName := fmt.Sprintf("%s_%s.txt", strings.ToLower(r.Zodiac.Name), r.Year)
	fullPath := filepath.Join(dirPath, fileName)

	// Create directory
	if err := os.MkdirAll(dirPath, 0755); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	// Write file
	if err := os.WriteFile(fullPath, []byte(content), 0644); err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	fmt.Println("✅ Saved reading to", fullPath)
	return nil
}
