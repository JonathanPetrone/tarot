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
// Fixed to handle case sensitivity and match your actual file structure
func getQualityAgentFilePath(year, month string) string {
	basePath, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting working directory:", err)
		return ""
	}

	// Try different case combinations to find the file
	possiblePaths := []string{
		filepath.Join(basePath, "QualityAgent", year, strings.Title(month), fmt.Sprintf("%s_review.json", strings.Title(month))),
		filepath.Join(basePath, "QualityAgent", year, month, fmt.Sprintf("%s_review.json", month)),
		filepath.Join(basePath, "QualityAgent", year, strings.ToUpper(month), fmt.Sprintf("%s_review.json", strings.ToUpper(month))),
	}

	for _, path := range possiblePaths {
		if _, err := os.Stat(path); err == nil {
			return path
		}
	}

	// Return the first path as fallback
	return possiblePaths[0]
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
// Fixed to properly handle empty summaries and complex formatting
func parseSignUpdates(content string) []SignUpdate {
	var updates []SignUpdate

	// Split by "Sign: " and process each section
	sections := strings.Split(content, "Sign: ")

	for _, section := range sections {
		section = strings.TrimSpace(section)
		if section == "" {
			continue
		}

		// Find the sign name (first line)
		lines := strings.Split(section, "\n")
		if len(lines) == 0 {
			continue
		}

		signName := strings.TrimSpace(lines[0])
		if signName == "" {
			continue
		}

		// Join the rest of the content
		restContent := strings.Join(lines[1:], "\n")

		var summary, finalWhisper string

		// Find the indices of "Summary:" and "Final Whisper:"
		summaryIndex := strings.Index(restContent, "Summary:")
		finalWhisperIndex := strings.Index(restContent, "Final Whisper:")

		if summaryIndex >= 0 && finalWhisperIndex >= 0 {
			// Both exist - extract summary between "Summary:" and "Final Whisper:"
			summaryStart := summaryIndex + len("Summary:")
			summaryContent := restContent[summaryStart:finalWhisperIndex]
			summary = strings.TrimSpace(summaryContent)

			// Extract final whisper after "Final Whisper:"
			finalWhisperStart := finalWhisperIndex + len("Final Whisper:")
			finalWhisperContent := restContent[finalWhisperStart:]
			finalWhisper = strings.TrimSpace(finalWhisperContent)

		} else if finalWhisperIndex >= 0 {
			// Only final whisper exists
			finalWhisperStart := finalWhisperIndex + len("Final Whisper:")
			finalWhisperContent := restContent[finalWhisperStart:]
			finalWhisper = strings.TrimSpace(finalWhisperContent)

		} else if summaryIndex >= 0 {
			// Only summary exists
			summaryStart := summaryIndex + len("Summary:")
			summaryContent := restContent[summaryStart:]
			summary = strings.TrimSpace(summaryContent)
		}

		// Clean up extra whitespace
		summary = regexp.MustCompile(`\s+`).ReplaceAllString(summary, " ")
		finalWhisper = regexp.MustCompile(`\s+`).ReplaceAllString(finalWhisper, " ")
		summary = strings.TrimSpace(summary)
		finalWhisper = strings.TrimSpace(finalWhisper)

		// Debug output
		log.Printf("Parsed sign: %s", signName)
		log.Printf("  Summary: '%s'", summary)
		log.Printf("  Final Whisper: '%s'", finalWhisper)

		// Add to updates if we have content
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

	// Extract summary - look for the paragraph with id="summary"
	summaryRegex := regexp.MustCompile(`<p[^>]*id="summary"[^>]*>(.*?)</p>`)
	summaryMatch := summaryRegex.FindStringSubmatch(htmlContent)
	if len(summaryMatch) > 1 {
		// Decode HTML entities and clean up
		summary = strings.ReplaceAll(summaryMatch[1], "&#39;", "'")
		summary = strings.ReplaceAll(summary, "&quot;", "\"")
		summary = strings.ReplaceAll(summary, "&amp;", "&")
		summary = strings.TrimSpace(summary)
	}

	// Extract final whisper - look for the paragraph with id="final_whisper"
	whisperRegex := regexp.MustCompile(`<p[^>]*id="final_whisper"[^>]*>(.*?)</p>`)
	whisperMatch := whisperRegex.FindStringSubmatch(htmlContent)
	if len(whisperMatch) > 1 {
		// Decode HTML entities and clean up
		finalWhisper = strings.ReplaceAll(whisperMatch[1], "&#39;", "'")
		finalWhisper = strings.ReplaceAll(finalWhisper, "&quot;", "\"")
		finalWhisper = strings.ReplaceAll(finalWhisper, "&amp;", "&")
		finalWhisper = strings.TrimSpace(finalWhisper)
	}

	return summary, finalWhisper, nil
}
func updateHTMLContent(filePath, newSummary, newFinalWhisper string) error {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("failed to read HTML file: %w", err)
	}

	htmlContent := string(content)

	// Update summary if provided
	if newSummary != "" {
		escapedSummary := strings.ReplaceAll(newSummary, "'", "&#39;")
		escapedSummary = strings.ReplaceAll(escapedSummary, "\"", "&quot;")

		// Fixed regex - more flexible attribute matching
		summaryRegex := regexp.MustCompile(`(?s)(<p[^>]*\bid="summary"[^>]*>)(.*?)(</p>)`)
		if summaryRegex.MatchString(htmlContent) {
			htmlContent = summaryRegex.ReplaceAllString(htmlContent, "${1}"+escapedSummary+"${3}")
			log.Printf("✅ Updated summary in %s", filePath)
		} else {
			log.Printf("❌ Summary regex didn't match in %s", filePath)
		}
	}

	// Update final whisper if provided
	if newFinalWhisper != "" {
		escapedWhisper := strings.ReplaceAll(newFinalWhisper, "'", "&#39;")
		escapedWhisper = strings.ReplaceAll(escapedWhisper, "\"", "&quot;")

		// Fixed regex for final whisper too
		whisperRegex := regexp.MustCompile(`(?s)(<p[^>]*\bid="final_whisper"[^>]*>)(.*?)(</p>)`)
		if whisperRegex.MatchString(htmlContent) {
			htmlContent = whisperRegex.ReplaceAllString(htmlContent, "${1}"+escapedWhisper+"${3}")
			log.Printf("✅ Updated final whisper in %s", filePath)
		} else {
			log.Printf("❌ Final whisper regex didn't match in %s", filePath)
		}
	}

	return os.WriteFile(filePath, []byte(htmlContent), 0644)
}

// UpdateHTMLFromQualityAgent updates HTML files with quality agent improvements
func UpdateHTMLFromQualityAgent(year, month string) error {
	qualityAgentPath := getQualityAgentFilePath(year, month)

	log.Printf("Looking for quality agent file at: %s", qualityAgentPath)

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

	if len(updates) == 0 {
		log.Printf("No updates found in quality agent file")
		return nil
	}

	// Process each sign update
	for _, update := range updates {
		htmlFilePath := getExistingHTMLFilePath(update.Sign, year, month)

		log.Printf("Processing %s: %s", update.Sign, htmlFilePath)

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
			log.Printf("  Old: %s", existingSummary[:min(50, len(existingSummary))]+"...")
			log.Printf("  New: %s", newSummary[:min(50, len(newSummary))]+"...")
		}

		if update.FinalWhisper != "" && update.FinalWhisper != existingWhisper {
			needsUpdate = true
			newWhisper = update.FinalWhisper
			log.Printf("📝 Final whisper update needed for %s", update.Sign)
			log.Printf("  Old: %s", existingWhisper[:min(50, len(existingWhisper))]+"...")
			log.Printf("  New: %s", newWhisper[:min(50, len(newWhisper))]+"...")
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

// Helper function for min
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
