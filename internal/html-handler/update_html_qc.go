package htmlhandler

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// QualityAgentResponse represents the structure of the quality agent output
type QualityAgentResponse struct {
	Choices []struct {
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
}

// SignUpdate represents the summary and final whisper for a zodiac sign
type SignUpdate struct {
	Sign         string
	Summary      string
	FinalWhisper string
}

// getQualityAgentFilePath builds the path to the quality agent output file
func getQualityAgentFilePath(year, month string) string {
	basePath, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting working directory:", err)
		return ""
	}

	// Assuming the quality agent output is in QualityAgent folder
	filePath := filepath.Join(basePath, "QualityAgent", year, month, fmt.Sprintf("%s_review.json", month))
	return filePath
}

// parseQualityAgentResponse extracts content from the quality agent JSON and parses it into sign updates
func parseQualityAgentResponse(filePath string) ([]SignUpdate, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read quality agent file: %w", err)
	}

	var resp QualityAgentResponse
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, fmt.Errorf("failed to parse quality agent JSON: %w", err)
	}

	if len(resp.Choices) == 0 {
		return nil, fmt.Errorf("no choices in quality agent JSON")
	}

	content := resp.Choices[0].Message.Content
	return parseSignUpdates(content), nil
}

// parseSignUpdates parses the content string into individual sign updates
func parseSignUpdates(content string) []SignUpdate {
	var updates []SignUpdate

	// Split content by "Sign:" to get individual sign sections
	sections := strings.Split(content, "Sign: ")

	for _, section := range sections {
		if strings.TrimSpace(section) == "" {
			continue
		}

		lines := strings.Split(section, "\n")
		if len(lines) == 0 {
			continue
		}

		// First line should contain the sign name
		signName := strings.TrimSpace(lines[0])
		if signName == "" {
			continue
		}

		var summary, finalWhisper string
		var currentSection string
		var currentText strings.Builder

		for i := 1; i < len(lines); i++ {
			line := strings.TrimSpace(lines[i])

			if line == "Summary:" {
				// Save previous section if exists
				if currentSection == "Final Whisper:" {
					finalWhisper = strings.TrimSpace(currentText.String())
				}
				currentSection = "Summary:"
				currentText.Reset()
			} else if line == "Final Whisper:" {
				// Save summary section
				if currentSection == "Summary:" {
					summary = strings.TrimSpace(currentText.String())
				}
				currentSection = "Final Whisper:"
				currentText.Reset()
			} else if line != "" {
				// Add content to current section
				if currentText.Len() > 0 {
					currentText.WriteString(" ")
				}
				currentText.WriteString(line)
			}
		}

		// Save the last section
		if currentSection == "Final Whisper:" {
			finalWhisper = strings.TrimSpace(currentText.String())
		} else if currentSection == "Summary:" {
			summary = strings.TrimSpace(currentText.String())
		}

		// Only add if we have valid content
		if signName != "" && (summary != "" || finalWhisper != "") {
			updates = append(updates, SignUpdate{
				Sign:         signName,
				Summary:      summary,
				FinalWhisper: finalWhisper,
			})
		}
	}

	return updates
}

// getExistingHTMLFilePath builds the path to the existing HTML file
func getExistingHTMLFilePath(sign, year, month string) string {
	outputFile := fmt.Sprintf("%s_%s_%s.html", strings.ToLower(sign), year, strings.ToLower(month))
	return filepath.Join("templates", "readings", year, month, outputFile)
}

// extractExistingSummaryAndWhisper extracts current summary and final whisper from HTML file
func extractExistingSummaryAndWhisper(filePath string) (summary, finalWhisper string, err error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return "", "", fmt.Errorf("failed to read HTML file: %w", err)
	}

	htmlContent := string(content)

	// Extract summary - look for the paragraph after the h2 title
	summaryRegex := regexp.MustCompile(`<h2[^>]*>Monthly Reading for[^<]*</h2>\s*<p[^>]*class="text-white[^"]*"[^>]*>(.*?)</p>`)
	summaryMatch := summaryRegex.FindStringSubmatch(htmlContent)
	if len(summaryMatch) > 1 {
		// Decode HTML entities and clean up
		summary = strings.ReplaceAll(summaryMatch[1], "&#39;", "'")
		summary = strings.ReplaceAll(summary, "&quot;", "\"")
		summary = strings.TrimSpace(summary)
	}

	// Extract final whisper - look for content after "Madame's Final Whisper"
	whisperRegex := regexp.MustCompile(`<h3[^>]*>🌟 Madame's Final Whisper</h3>\s*<p[^>]*>(.*?)</p>`)
	whisperMatch := whisperRegex.FindStringSubmatch(htmlContent)
	if len(whisperMatch) > 1 {
		// Decode HTML entities and clean up
		finalWhisper = strings.ReplaceAll(whisperMatch[1], "&#39;", "'")
		finalWhisper = strings.ReplaceAll(finalWhisper, "&quot;", "\"")
		finalWhisper = strings.TrimSpace(finalWhisper)
	}

	return summary, finalWhisper, nil
}

// updateHTMLContent updates the HTML content with new summary and/or final whisper
func updateHTMLContent(filePath, newSummary, newFinalWhisper string) error {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("failed to read HTML file: %w", err)
	}

	htmlContent := string(content)

	// Update summary if provided
	if newSummary != "" {
		// Escape quotes for HTML
		escapedSummary := strings.ReplaceAll(newSummary, "'", "&#39;")
		escapedSummary = strings.ReplaceAll(escapedSummary, "\"", "&quot;")

		summaryRegex := regexp.MustCompile(`(<h2[^>]*>Monthly Reading for[^<]*</h2>\s*<p[^>]*class="text-white[^"]*"[^>]*>)(.*?)(</p>)`)
		htmlContent = summaryRegex.ReplaceAllString(htmlContent, "${1}"+escapedSummary+"${3}")
	}

	// Update final whisper if provided
	if newFinalWhisper != "" {
		// Escape quotes for HTML
		escapedWhisper := strings.ReplaceAll(newFinalWhisper, "'", "&#39;")
		escapedWhisper = strings.ReplaceAll(escapedWhisper, "\"", "&quot;")

		whisperRegex := regexp.MustCompile(`(<h3[^>]*>🌟 Madame's Final Whisper</h3>\s*<p[^>]*>)(.*?)(</p>)`)
		htmlContent = whisperRegex.ReplaceAllString(htmlContent, "${1}"+escapedWhisper+"${3}")
	}

	// Write updated content back to file
	return os.WriteFile(filePath, []byte(htmlContent), 0644)
}

// UpdateHTMLFromQualityAgent updates HTML files with quality agent improvements
func UpdateHTMLFromQualityAgent(year, month string) error {
	qualityAgentPath := getQualityAgentFilePath(year, month)

	// Check if quality agent file exists
	if _, err := os.Stat(qualityAgentPath); os.IsNotExist(err) {
		return fmt.Errorf("quality agent file not found: %s", qualityAgentPath)
	}

	// Parse quality agent response
	updates, err := parseQualityAgentResponse(qualityAgentPath)
	if err != nil {
		return fmt.Errorf("failed to parse quality agent response: %w", err)
	}

	log.Printf("Found %d sign updates in quality agent file", len(updates))

	// Process each sign update
	for _, update := range updates {
		htmlFilePath := getExistingHTMLFilePath(update.Sign, year, month)

		// Check if HTML file exists
		if _, err := os.Stat(htmlFilePath); os.IsNotExist(err) {
			log.Printf("⚠️ HTML file not found for %s: %s", update.Sign, htmlFilePath)
			continue
		}

		// Extract existing content
		existingSummary, existingWhisper, err := extractExistingSummaryAndWhisper(htmlFilePath)
		if err != nil {
			log.Printf("⚠️ Failed to extract existing content for %s: %v", update.Sign, err)
			continue
		}

		// Check if updates are needed
		needsUpdate := false
		var newSummary, newWhisper string

		if update.Summary != "" && update.Summary != existingSummary {
			needsUpdate = true
			newSummary = update.Summary
			log.Printf("📝 Summary update needed for %s", update.Sign)
		}

		if update.FinalWhisper != "" && update.FinalWhisper != existingWhisper {
			needsUpdate = true
			newWhisper = update.FinalWhisper
			log.Printf("📝 Final whisper update needed for %s", update.Sign)
		}

		// Apply updates if needed
		if needsUpdate {
			err := updateHTMLContent(htmlFilePath, newSummary, newWhisper)
			if err != nil {
				log.Printf("❌ Failed to update HTML for %s: %v", update.Sign, err)
				continue
			}
			log.Printf("✅ Successfully updated HTML for %s", update.Sign)
		} else {
			log.Printf("ℹ️ No updates needed for %s", update.Sign)
		}
	}

	return nil
}
