package htmlhandler

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
)

type Statistics struct {
	MajorArcana    int
	MinorArcana    int
	MostCommonSuit []string
	MostCommonRank []string
}

func ParseStatistics(filePath string) (*Statistics, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	stats := Statistics{}
	suitCount := map[string]int{}
	rankCount := map[string]int{}

	suitRegex := regexp.MustCompile(`^(Cups|Pentacles|Swords|Wands):\s+(\d+)`)
	rankRegex := regexp.MustCompile(`^(Twos|Threes|Fours|Fives|Sixes|Sevens|Eights|Nines|Tens|Pages|Knights|Queens|Kings|Aces):\s+(\d+)`)
	arcanaRegex := regexp.MustCompile(`^(Major|Minor) Arcana Cards:\s+(\d+)`)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if matches := suitRegex.FindStringSubmatch(line); matches != nil {
			count, _ := strconv.Atoi(matches[2])
			suitCount[matches[1]] = count
		}

		if matches := rankRegex.FindStringSubmatch(line); matches != nil {
			count, _ := strconv.Atoi(matches[2])
			rankCount[matches[1]] = count
		}

		if matches := arcanaRegex.FindStringSubmatch(line); matches != nil {
			count, _ := strconv.Atoi(matches[2])
			if matches[1] == "Major" {
				stats.MajorArcana = count
			} else {
				stats.MinorArcana = count
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	// Find most common suit
	maxSuit := 0
	for suit, count := range suitCount {
		if count > maxSuit {
			maxSuit = count
			stats.MostCommonSuit = []string{suit}
		} else if count == maxSuit {
			stats.MostCommonSuit = append(stats.MostCommonSuit, suit)
		}
	}

	// Find most common rank
	maxRank := 0
	for rank, count := range rankCount {
		if count > maxRank {
			maxRank = count
			stats.MostCommonRank = []string{rank}
		} else if count == maxRank && count > 0 {
			stats.MostCommonRank = append(stats.MostCommonRank, rank)
		}
	}

	// If no rank reaches a count of 2, write a fitting message
	if maxRank < 2 {
		stats.MostCommonRank = []string{"Mixed Ranks"}
	}

	return &stats, nil
}
