package main

import (
	"fmt"
	"strings"

	htmlhandler "github.com/jonathanpetrone/aitarot/internal/html-handler"
)

func main() {
	stats, err := htmlhandler.ParseStatistics("monthlyreadings/2025/july/sagittarius_2025.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("Major Arcana Cards: %d\n", stats.MajorArcana)
	fmt.Printf("Minor Arcana Cards: %d\n", stats.MinorArcana)
	fmt.Printf("Most Common Suite: %s\n", strings.Join(stats.MostCommonSuit, ", "))
	fmt.Printf("Most Common Rank: %s\n", strings.Join(stats.MostCommonRank, ", "))
}
