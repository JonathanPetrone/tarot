package main

import (
	"github.com/jonathanpetrone/aitarot/internal/astrology"
	"github.com/jonathanpetrone/aitarot/internal/tarot"
)

// rand.Seed() If I want predicability for testing

func main() {
	aries := astrology.StarSigns[0]
	spread := tarot.ReadSpread(tarot.CelticCross)
	spread2 := tarot.ReadSpread(tarot.PastPresentFuture)

	tarot.FormatReading(tarot.CelticCross, spread, aries)
	tarot.FormatReading(tarot.PastPresentFuture, spread2, aries)

	stats := tarot.Stats{}
	tarot.AnalyzeSpreadTarot(spread, &stats)
	stats.Print()
}
