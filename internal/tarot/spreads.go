package tarot

import (
	"fmt"

	"github.com/jonathanpetrone/aitarot/internal/astrology"
)

type Spread struct {
	Name     string
	Length   int
	Template []SpreadCard
}

type SpreadCard struct {
	Position int
	Context  string
	Card     TarotCard
}

var PastPresentFuture = Spread{
	Name:   "Past-Present-Future",
	Length: 3,
	Template: []SpreadCard{
		{Position: 1, Context: "Past"},
		{Position: 2, Context: "Present"},
		{Position: 3, Context: "Future"},
	},
}

var CelticCross = Spread{
	Name:   "Celtic Cross",
	Length: 10,
	Template: []SpreadCard{
		{Position: 1, Context: "Current Situation"},
		{Position: 2, Context: "Challenges"},
		{Position: 3, Context: "Strengths"},
		{Position: 4, Context: "What to Focus On"},
		{Position: 5, Context: "Past Energy"},
		{Position: 6, Context: "Near Future"},
		{Position: 7, Context: "Advice"},
		{Position: 8, Context: "External Influences"},
		{Position: 9, Context: "Hopes and Fears"},
		{Position: 10, Context: "Likely Outcome"},
	},
}

func ReadSpread(s Spread) []SpreadCard {
	reading := make([]SpreadCard, s.Length)
	copy(reading, s.Template)

	drawnCards := DrawCards(s.Length)

	for i, j := range drawnCards {
		reading[i].Card = j
	}

	return reading
}

func FormatReading(spread Spread, reading []SpreadCard, ss astrology.StarSign, aiOutput bool) {
	starSignString := ss.Name

	if ss.Name == "" {
		fmt.Println("The recipient has not provided their star sign.")
		starSignString = "(You did not provide a star sign)"
	}

	fmt.Println("----------------------------------------------------------")
	fmt.Printf("Here is your %s reading, %s\n", spread.Name, starSignString)
	fmt.Println("----------------------------------------------------------")

	for _, position := range reading {
		fmt.Printf("%2d. %-35s -> %s\n", position.Position, position.Context, position.Card.Name)
	}

	fmt.Println("")
}
