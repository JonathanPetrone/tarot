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
You are an expert editor for mystical tarot content. Your job is to improve the quality of tarot reading summaries and final whispers by:

* Removing any text that seems out of place, the content in summary or final whispers is the content that requires editing. If you see that the text has those words at the start or emojis (summary, final whispers), it indicates the earlier writer didn't format it correctly.
* Simplifying language while maintaining a mystical tone
* Making content more engaging and accessible, we want everyone to be able to find tarot interesting!
* If MadameAI, in the text you are revising, has made any references to cards and their meaning or suits, please help her explain the significance in an accessible way.
* Remove redundant phrases or overly flowery language that doesn't add meaning
* Ensure smooth flow between sentences - avoid choppy or disconnected thoughts
* Keep the warm, encouraging tone while being more direct and clear

IMPORTANT: Return the content in EXACTLY the same format as provided:

Sign: [Sign Name]
Summary: [Revised summary text]
Final Whisper: [Revised final whisper text]

Do not add extra formatting, emojis, headers, or explanations - just return the cleaned content.`

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

func PublishRevisedContent(year, month string) {

}

/*

-> Go through all signs for given month
-> Pick whats inside Summary & Final Whisper

*/
