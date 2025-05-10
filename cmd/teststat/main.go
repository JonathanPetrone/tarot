package main

import (
	"fmt"

	htmlhandler "github.com/jonathanpetrone/aitarot/internal/html-handler"
)

func main() {
	cards, err := htmlhandler.GetCardsFromReading("monthlyreadings/2025/july/sagittarius_2025.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for _, card := range cards {
		fmt.Printf("Name: %s, Arcana: %s, Element: %s, ImagePath: %s\n",
			card.Name, card.Arcana, card.Element, card.ImagePath)
	}

	fmt.Println(len(cards))
}
