package tarot

import (
	"fmt"
	"strings"

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

func FormatReading(spread Spread, reading []SpreadCard, ss astrology.StarSign, aiOutput bool) string {
	var b strings.Builder

	b.WriteString(`<div class="p-8">`)
	b.WriteString(fmt.Sprintf(`<h2 class="text-white text-3xl mb-4">%s Reading</h2>`, ss.Name))

	for _, position := range reading {
		b.WriteString(fmt.Sprintf(`<p class="text-white">%2d. %-35s -> %s</p>`, position.Position, position.Context, position.Card.Name))
	}

	b.WriteString(`</div>`)
	return b.String()
}
