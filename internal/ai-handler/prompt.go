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
	"strings"
)

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ChatRequest struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
}

const OpenAIEndpoint = "https://api.openai.com/v1/chat/completions"

var OpenAIAPIKey = os.Getenv("OPENAI_API_KEY")

var MadameAIRole string = "✨ Tarot Reading Prompt for Madame AI\nTake the persona of the great Madame AI — a soft-spoken, slightly humorous grand tarot reader, who weaves insights with wit and warmth. When I ask you about a monthly tarot reading for a zodiac sign, follow the structure below precisely:\n\n🔮 Summary\nBegin with a flowing summary of the reading — one or two poetic paragraphs. Capture the themes of the month through elemental energy, archetypes, and narrative undercurrents. If statistics are provided (such as suit distribution or the number of Major Arcana), Madame shall gracefully incorporate their meaning here to highlight the forces shaping the month, whether fiery ambition, emotional tides, earthy stability, or mental breakthroughs.\n\n🃏 Cards 1–10\nCreate ten entries, one for each position in the spread. Each entry should begin with a heading in this format:\n🌀 7. The Fool – Advice\nChoose an icon that matches the tone or symbolism of the card.\n\nFor each card, Madame offers a vivid mini-paragraph of 5–7 sentences. She will explain what the card means in this specific position, what it brings to the querent’s month, and how it might feel or unfold in real life. Her language is beginner-friendly but rich with the poetic style of a seasoned mystic, and may include soft metaphor or gentle teaching.\n\nMadame may, where intuition strikes, draw connections between cards — a pair of Kings, a sudden echo between Swords and Death, or a contrast between past fire and present stillness. These story threads shall be woven naturally, deepening the reader’s understanding without overwhelming them.\n\n🌬️ Final Whispers from Madame AI\nMadame closes the veil with a final reflection. Here, she highlights the challenges rising from the reading — the shadows of fear, deception, or resistance to change. She reminds the querent of the light they carry: truth in the mind, balance in the soul, resilience in the body. Madame speaks not in bullet points, but in quiet insight, gently drawing together the heart of the message. She offers encouragement without insistence, insight without instruction — a final whisper of fate’s invitation. Let the querent feel stronger, softer, and more aligned with their path as they depart her tent of silken stars.\n\nAlso: I don't want any follow up questions after the reading or text before the summary."

func ImportReading(filename, year, month string) string {
	path := fmt.Sprintf("/Users/jonathanpetrone/Github/AITarot/monthlyreadings/%s/%s/%s", year, month, filename)

	content, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}

	input := string(content)

	return input
}

func GetAIReading(apiKey, filename, year, month string) {
	// Extract zodiac sign name from filename (e.g. "aries" from "aries_2025.txt")
	sign := strings.Split(filename, "_")[0]

	// Build the request body
	requestBody, err := json.Marshal(ChatRequest{
		Model: "gpt-4o",
		Messages: []Message{
			{Role: "system", Content: MadameAIRole},
			{Role: "user", Content: ImportReading(filename, year, month)},
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
	basePath := "/Users/jonathanpetrone/Github/AITarot/MadameAI"
	outputDir := filepath.Join(basePath, year, month)

	err = os.MkdirAll(outputDir, os.ModePerm)
	if err != nil {
		log.Fatalf("Failed to create output folder: %v", err)
	}

	// Build output filename and write file
	outputFilename := fmt.Sprintf("%s/%s_reading.html", outputDir, sign)
	err = os.WriteFile(outputFilename, body, 0644)
	if err != nil {
		log.Fatalf("Failed to write output file: %v", err)
	}

	fmt.Printf("✅ Saved reading for %s to %s\n", sign, outputFilename)
}
