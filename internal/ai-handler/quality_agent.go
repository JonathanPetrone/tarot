package aihandler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/jonathanpetrone/aitarot/internal/astrology"
)

type ReadingContent struct {
	Sign         string
	Summary      string
	FinalWhisper string
	FilePath     string
}

type ReadingContents struct {
	Contents []ReadingContent
}

var QualityAIRole string = `# Your role
You are an expert editor for mystical tarot content. Your job is to REFINE and ENHANCE tarot reading summaries and final whispers while preserving their mystical essence and tarot-specific content.

## PRESERVE these elements:
* ALL references to specific tarot cards (The Star, The Emperor, Nine of Wands, etc.)
* Suit mentions (Swords, Cups, Pentacles, Wands) and their meanings
* Mystical and astrological language that creates atmosphere
* Unique personality and voice for each zodiac sign
* The overall length and depth of the original content

## IMPROVE these elements:
* Fix awkward phrasing or unclear sentences
* Remove only truly redundant phrases (like repeating the same idea twice in one paragraph)
* Ensure smooth transitions between sentences
* Make card references more accessible by briefly explaining their significance when helpful
* Fix any grammar or punctuation issues
* Remove obvious formatting errors (like emoji symbols appearing in text, or "Summary:" appearing at the start of content)

## CRITICAL REQUIREMENTS:
* NEVER leave a summary empty - if the original summary is missing or very short, create one based on themes from the final whisper and card meanings
* REMOVE ALL EMOJIS from the text content (🔮, 🌙, etc.) - the text should be pure prose
* Each summary should be 3-6 sentences long minimum
* Each final whisper should be 3-6 sentences long

## TONE GUIDELINES:
* Maintain the warm, encouraging, and mystical tone
* Keep the poetic language that makes tarot readings special
* Each sign should retain its distinct personality
* Content should feel magical and insightful, not generic

## CRITICAL: 
* If a summary is substantial and well-written, make minimal changes
* If a summary is missing or very short, you MUST create one using the card themes mentioned in the reading
* NEVER remove all the content and replace it with generic advice
* NEVER make all the signs sound the same
* NO EMOJIS anywhere in the output text

Return the content in EXACTLY this format:

Sign: [Sign Name]
Summary: [Complete summary text - never empty, 3-6 sentences, no emojis]
Final Whisper: [Complete final whisper text - 3-6 sentences, no emojis]

Focus on polish and clarity while ensuring every sign has both a meaningful summary and final whisper, completely free of emoji symbols.`

func ExtractReadingFromHTML(filePath string) (*ReadingContent, error) {
	// Read the HTML file
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	content := string(data)

	// Create regex patterns for summary and final whisper
	summaryRegex := regexp.MustCompile(`<p id="summary"[^>]*>(.*?)</p>`)
	finalWhisperRegex := regexp.MustCompile(`<p id="final_whisper"[^>]*>(.*?)</p>`)

	reading := &ReadingContent{
		FilePath: filePath,
	}

	// Extract summary
	if matches := summaryRegex.FindStringSubmatch(content); len(matches) > 1 {
		reading.Summary = cleanHTMLContent(matches[1])
	}

	// Extract final whisper
	if matches := finalWhisperRegex.FindStringSubmatch(content); len(matches) > 1 {
		reading.FinalWhisper = cleanHTMLContent(matches[1])
	}

	// Extract sign from filename (e.g., "aquarius_2025_may.html" -> "Aquarius")
	filename := filepath.Base(filePath)
	if parts := strings.Split(filename, "_"); len(parts) > 0 {
		reading.Sign = strings.Title(parts[0])
	}

	return reading, nil
}

func cleanHTMLContent(text string) string {
	// Remove HTML entities
	text = strings.ReplaceAll(text, "&#39;", "'")
	text = strings.ReplaceAll(text, "&quot;", "\"")
	text = strings.ReplaceAll(text, "&amp;", "&")

	// Remove extra whitespace
	text = regexp.MustCompile(`\s+`).ReplaceAllString(text, " ")

	return strings.TrimSpace(text)
}

func TestExtraction() {
	reading, err := ExtractReadingFromHTML("templates/readings/2025/may/aquarius_2025_may.html")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Sign: %s\n", reading.Sign)
	fmt.Printf("Summary: %s\n", reading.Summary)
	fmt.Printf("Final Whisper: %s\n", reading.FinalWhisper)
}

func PrepareQualityReview(year, month string) *ReadingContents {
	// Create the container struct
	allReadings := &ReadingContents{
		Contents: make([]ReadingContent, 0, len(astrology.ZodiacSigns)),
	}

	for _, sign := range astrology.ZodiacSigns {
		filepath := fmt.Sprintf("templates/readings/%s/%s/%s_%s_%s.html", year, month, strings.ToLower(sign.Name), year, month)

		reading, err := ExtractReadingFromHTML(filepath)
		if err != nil {
			log.Printf("Failed to extract reading for %s: %v", sign.Name, err)
			continue // Skip this sign and continue with others
		}

		// Dereference the pointer and append to slice
		allReadings.Contents = append(allReadings.Contents, *reading)
	}

	return allReadings
}

func GetQualityReview(apiKey, year, month string) {
	allReadings := PrepareQualityReview(year, month)
	content := ""

	for _, reading := range allReadings.Contents {
		content += fmt.Sprintf("Sign: %s\nSummary: %s\nFinal Whisper: %s\n\n",
			reading.Sign, reading.Summary, reading.FinalWhisper)
	}

	// Build the request body
	requestBody, err := json.Marshal(ChatRequest{
		Model: "gpt-4o",
		Messages: []Message{
			{Role: "system", Content: QualityAIRole}, // Note: you had MadameAIRole but declared QualityAIRole
			{Role: "user", Content: content},
		},
	})
	if err != nil {
		log.Fatalf("Failed to marshal request: %v", err)
	}

	// Create HTTP request
	req, err := http.NewRequest("POST", OpenAIEndpoint, bytes.NewBuffer(requestBody))
	if err != nil {
		log.Fatalf("Failed to create request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	// Send request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Request failed: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read response: %v", err)
	}

	// Create output directory path
	basePath := "./QualityAgent"
	outputDir := filepath.Join(basePath, year, month)

	err = os.MkdirAll(outputDir, os.ModePerm)
	if err != nil {
		log.Fatalf("Failed to create output folder: %v", err)
	}

	// Build output filename and write file
	outputFilename := fmt.Sprintf("%s/%s_review.json", outputDir, month)
	err = os.WriteFile(outputFilename, body, 0644)
	if err != nil {
		log.Fatalf("Failed to write output file: %v", err)
	}

	fmt.Printf("✅ Output saved in %s\n", outputFilename)

}
