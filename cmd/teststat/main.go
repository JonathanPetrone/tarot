package main

import (
	"fmt"

	"github.com/jonathanpetrone/aitarot/internal/tarot"
)

func main() {
	// admin.CreateNewReadings("Aries", "2025", "December")
	cards := tarot.DrawCards(1)
	fmt.Println(cards)
}
