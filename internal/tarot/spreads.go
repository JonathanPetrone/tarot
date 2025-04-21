package tarot

import (
	"fmt"

	"github.com/jonathanpetrone/aitarot/internal/astrology"
)

type Spread struct {
	Position int
	Context  string
	Card     TarotCard
}

var CelticCross = []Spread{
	{Position: 1, Context: "Current Situation: "},
	{Position: 2, Context: "Challenges/Obstacles: "},
	{Position: 3, Context: "Strengths: "},
	{Position: 4, Context: "Need to Focus on: "},
	{Position: 5, Context: "The Past or Leaving Energy: "},
	{Position: 6, Context: "Near Future and Upcoming Energy: "},
	{Position: 7, Context: "Advice: "},
	{Position: 8, Context: "External influences: "},
	{Position: 9, Context: "Hopes and Fears: "},
	{Position: 10, Context: "Current Future Outcome: "},
}

func ReadCelticCross() []Spread {
	reading := make([]Spread, len(CelticCross))
	copy(reading, CelticCross)

	drawnCards := DrawCards(10)

	for i, j := range drawnCards {
		reading[i].Card = j
	}

	return reading
}

func FormatCelticCross(spread []Spread, ss astrology.StarSign) {
	starSignString := ss.Name

	if ss.Name == "" {
		fmt.Println("The recipient has not provided their star sign.")
		starSignString = "(You did not provide a star sign)"
	}

	fmt.Println("----------------------------------------------------------")
	fmt.Printf("Here is your Celtic Cross reading, %s\n", starSignString)
	fmt.Println("----------------------------------------------------------")

	for _, position := range spread {
		fmt.Printf("%2d. %-35s -> %s\n", position.Position, position.Context, position.Card.Name)
	}
}
