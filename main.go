package main

import (
	"github.com/jonathanpetrone/aitarot/internal/astrology"
	"github.com/jonathanpetrone/aitarot/internal/tarot"
)

// rand.Seed() If I want predicability for testing

func main() {
	// drawnCards := tarot.DrawCards(1)

	/*for _, card := range drawnCards {
		fmt.Println(card.Name)
	} */

	aries := astrology.StarSigns[0]
	spread := tarot.ReadCelticCross()

	tarot.FormatCelticCross(spread, aries)
}
