package aihandler

import (
	"os"
	"regexp"
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
}

func ParseMonthlyReading(filepath string) (MonthlyReading, error) {
	var reading MonthlyReading

	content, err := os.ReadFile(filepath)
	if err != nil {
		return reading, err
	}
	input := string(content)

	// Split the input into sections
	finalWhisperIndex := strings.Index(input, "Final whispers from Madame AI")
	mainPart := input
	finalWhisper := ""

	if finalWhisperIndex != -1 {
		mainPart = strings.TrimSpace(input[:finalWhisperIndex])
		finalWhisper = strings.TrimSpace(input[finalWhisperIndex:])
	}

	// First paragraph is the summary
	paragraphs := strings.SplitN(mainPart, "\n\n", 2)
	if len(paragraphs) >= 1 {
		reading.Summary = strings.TrimSpace(paragraphs[0])
	}

	// Find all cards
	cardPattern := regexp.MustCompile(`(?m)^Card \d+ - .+?:`)
	cardIndices := cardPattern.FindAllStringIndex(mainPart, -1)

	for i, indices := range cardIndices {
		start := indices[0]
		end := len(mainPart)
		if i+1 < len(cardIndices) {
			end = cardIndices[i+1][0]
		}
		cardBlock := strings.TrimSpace(mainPart[start:end])

		// Split the card block into title and description
		lines := strings.SplitN(cardBlock, "\n", 2)
		if len(lines) == 2 {
			title := strings.TrimSpace(lines[0])
			description := strings.TrimSpace(lines[1])

			reading.Cards = append(reading.Cards, Card{
				Title:       title,
				Description: description,
			})
		}
	}

	// Final Whispers
	reading.FinalWhispers = finalWhisper
	//fmt.Printf("%+v\n", reading)

	return reading, nil
}
