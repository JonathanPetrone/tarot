package htmlhandler

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

type MadameAIResponse struct {
	Choices []struct {
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
}

func ExtractContentFromResponse(sign, year, month string) string {
	filePath := fmt.Sprintf(
		"/Users/jonathanpetrone/Github/AITarot/MadameAI/%s/%s/%s_reading.html",
		year, strings.ToLower(month), strings.ToLower(sign),
	)

	data, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}

	var resp MadameAIResponse
	if err := json.Unmarshal(data, &resp); err != nil {
		log.Fatalf("Failed to parse JSON: %v", err)
	}

	if len(resp.Choices) == 0 {
		log.Fatal("No choices in JSON")
	}

	return resp.Choices[0].Message.Content
}

func splitFlexibleParagraphs(input string) []string {
	re := regexp.MustCompile(`\n{2,}`) // match 2 or more newlines
	return re.Split(strings.TrimSpace(input), -1)
}

func SplitMadameAIContent(content string) ([]string, error) {
	content = strings.Replace(content, "🔮 Summary\n", "", 1)
	content = strings.Replace(content, "🌬️ Final Whispers from Madame AI", "", 1)

	cardPattern := regexp.MustCompile(`(?m)^.*?\b(\d{1,2})\.\s`)
	indices := cardPattern.FindAllStringIndex(content, -1)

	if len(indices) < 10 {
		return nil, fmt.Errorf("expected 10 card sections, found %d", len(indices))
	}

	// Extract summary (everything before card 1)
	summary := strings.TrimSpace(content[:indices[0][0]])
	parts := []string{summary}

	// Extract cards 1–9
	for i := 0; i < 9; i++ {
		start := indices[i][0]
		end := indices[i+1][0]
		parts = append(parts, strings.TrimSpace(content[start:end]))
	}

	// Extract card 10 and final whisper separately
	card10Start := indices[9][0]
	card10End := card10Start + strings.Index(content[card10Start:], "\n\n\n")
	if card10End == -1 || card10End <= card10Start {
		// fallback if no double newline: card10 goes to end, no final whisper
		card10End = len(content)
	}

	card10 := strings.TrimSpace(content[card10Start:card10End])
	parts = append(parts, card10)

	finalWhisper := strings.TrimSpace(content[card10End:])
	parts = append(parts, finalWhisper)

	if len(parts) != 12 {
		return nil, fmt.Errorf("final split result had %d parts (expected 12)", len(parts))
	}

	for i := range parts {
		re := regexp.MustCompile(`\*\*(.*?)\*\*|__(.*?)__|~~(.*?)~~`)
		parts[i] = re.ReplaceAllString(parts[i], "")
	}

	return parts, nil
}
