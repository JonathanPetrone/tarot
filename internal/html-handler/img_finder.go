package htmlhandler

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/jonathanpetrone/aitarot/internal/tarot"
)

func GetCardsFromReading(filePath string) ([]tarot.TarotCard, error) {
	// Initialize the result slice
	sliceOfCards := []tarot.TarotCard{}

	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Improved regex for matching card names
	cardRegex := regexp.MustCompile(`(?m)^\s*\d+\.\s+.+?\s+->\s+(.+?)\s*$`)

	// Create a map for quick lookup
	cardMap := make(map[string]tarot.TarotCard)
	for _, card := range tarot.Deck {
		cardMap[strings.ToLower(card.Name)] = card
	}

	// Scan the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Printf("Line scanned: '%s'\n", line) // Debugging line

		if matches := cardRegex.FindStringSubmatch(line); matches != nil {
			cardName := strings.TrimSpace(matches[1])
			cardName = strings.ToLower(cardName)

			// Debugging output
			fmt.Printf("Looking for: '%s'\n", cardName)

			if card, exists := cardMap[cardName]; exists {
				fmt.Printf("✅ Appending Card: %s\n", card.Name)
				sliceOfCards = append(sliceOfCards, card)
			} /*else {
				fmt.Printf("⚠️ Card not found in deck: '%s'\n", cardName)
			}*/
		} else {
			fmt.Printf("🚫 No match for line: '%s'\n", line)
		}
	}

	// Return error if scanning failed
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	fmt.Printf("Total cards found: %d\n", len(sliceOfCards))
	return sliceOfCards, nil
}
