package tarot

import (
	"math/rand"
)

type TarotCard struct {
	Name     string
	Number   int
	Arcana   string
	Suite    string
	Zodiac   string
	Planet   string
	Element  string
	CardRole string
}

var Deck = []TarotCard{
	// Major Arcana
	{"The Fool", 0, "Major", "", "", "Uranus", "Air", ""},
	{"The Magician", 1, "Major", "", "", "Mercury", "Air", ""},
	{"The High Priestess", 2, "Major", "", "", "Moon", "Water", ""},
	{"The Empress", 3, "Major", "", "", "Venus", "Earth", ""},
	{"The Emperor", 4, "Major", "", "Aries", "", "Fire", ""},
	{"The Hierophant", 5, "Major", "", "Taurus", "", "Earth", ""},
	{"The Lovers", 6, "Major", "", "Gemini", "", "Air", ""},
	{"The Chariot", 7, "Major", "", "Cancer", "", "Water", ""},
	{"Strength", 8, "Major", "", "Leo", "", "Fire", ""},
	{"The Hermit", 9, "Major", "", "Virgo", "", "Earth", ""},
	{"Wheel of Fortune", 10, "Major", "", "", "Jupiter", "Fire", ""},
	{"Justice", 11, "Major", "", "Libra", "", "Air", ""},
	{"The Hanged Man", 12, "Major", "", "", "Neptune", "Water", ""},
	{"Death", 13, "Major", "", "Scorpio", "", "Water", ""},
	{"Temperance", 14, "Major", "", "Sagittarius", "", "Fire", ""},
	{"The Devil", 15, "Major", "", "Capricorn", "", "Earth", ""},
	{"The Tower", 16, "Major", "", "", "Mars", "Fire", ""},
	{"The Star", 17, "Major", "", "Aquarius", "", "Air", ""},
	{"The Moon", 18, "Major", "", "Pisces", "", "Water", ""},
	{"The Sun", 19, "Major", "", "", "Sun", "Fire", ""},
	{"Judgement", 20, "Major", "", "", "Pluto", "Fire", ""},
	{"The World", 21, "Major", "", "", "Saturn", "Earth", ""},

	// Cups
	{"Ace of Cups", 1, "Minor", "Cups", "", "", "Water", ""},
	{"Two of Cups", 2, "Minor", "Cups", "Cancer", "Venus", "Water", ""},
	{"Three of Cups", 3, "Minor", "Cups", "Cancer", "Mercury", "Water", ""},
	{"Four of Cups", 4, "Minor", "Cups", "Cancer", "Moon", "Water", ""},
	{"Five of Cups", 5, "Minor", "Cups", "Scorpio", "Mars", "Water", ""},
	{"Six of Cups", 6, "Minor", "Cups", "Scorpio", "Sun", "Water", ""},
	{"Seven of Cups", 7, "Minor", "Cups", "Scorpio", "Venus", "Water", ""},
	{"Eight of Cups", 8, "Minor", "Cups", "Pisces", "Saturn", "Water", ""},
	{"Nine of Cups", 9, "Minor", "Cups", "Pisces", "Jupiter", "Water", ""},
	{"Ten of Cups", 10, "Minor", "Cups", "Pisces", "Mars", "Water", ""},
	{"Page of Cups", 11, "Minor", "Cups", "", "", "Water", "Page"},
	{"Knight of Cups", 12, "Minor", "Cups", "", "", "Water", "Knight"},
	{"Queen of Cups", 13, "Minor", "Cups", "", "", "Water", "Queen"},
	{"King of Cups", 14, "Minor", "Cups", "", "", "Water", "King"},

	// Pentacles
	{"Ace of Pentacles", 1, "Minor", "Pentacles", "", "", "Earth", ""},
	{"Two of Pentacles", 2, "Minor", "Pentacles", "Capricorn", "Jupiter", "Earth", ""},
	{"Three of Pentacles", 3, "Minor", "Pentacles", "Capricorn", "Mars", "Earth", ""},
	{"Four of Pentacles", 4, "Minor", "Pentacles", "Taurus", "Sun", "Earth", ""},
	{"Five of Pentacles", 5, "Minor", "Pentacles", "Taurus", "Mercury", "Earth", ""},
	{"Six of Pentacles", 6, "Minor", "Pentacles", "Taurus", "Moon", "Earth", ""},
	{"Seven of Pentacles", 7, "Minor", "Pentacles", "Virgo", "Saturn", "Earth", ""},
	{"Eight of Pentacles", 8, "Minor", "Pentacles", "Virgo", "Sun", "Earth", ""},
	{"Nine of Pentacles", 9, "Minor", "Pentacles", "Virgo", "Venus", "Earth", ""},
	{"Ten of Pentacles", 10, "Minor", "Pentacles", "Capricorn", "Mercury", "Earth", ""},
	{"Page of Pentacles", 11, "Minor", "Pentacles", "", "", "Earth", "Page"},
	{"Knight of Pentacles", 12, "Minor", "Pentacles", "", "", "Earth", "Knight"},
	{"Queen of Pentacles", 13, "Minor", "Pentacles", "", "", "Earth", "Queen"},
	{"King of Pentacles", 14, "Minor", "Pentacles", "", "", "Earth", "King"},

	// Swords
	{"Ace of Swords", 1, "Minor", "Swords", "", "", "Air", ""},
	{"Two of Swords", 2, "Minor", "Swords", "Libra", "Moon", "Air", ""},
	{"Three of Swords", 3, "Minor", "Swords", "Libra", "Saturn", "Air", ""},
	{"Four of Swords", 4, "Minor", "Swords", "Libra", "Jupiter", "Air", ""},
	{"Five of Swords", 5, "Minor", "Swords", "Aquarius", "Venus", "Air", ""},
	{"Six of Swords", 6, "Minor", "Swords", "Aquarius", "Mercury", "Air", ""},
	{"Seven of Swords", 7, "Minor", "Swords", "Aquarius", "Moon", "Air", ""},
	{"Eight of Swords", 8, "Minor", "Swords", "Gemini", "Jupiter", "Air", ""},
	{"Nine of Swords", 9, "Minor", "Swords", "Gemini", "Mars", "Air", ""},
	{"Ten of Swords", 10, "Minor", "Swords", "Gemini", "Sun", "Air", ""},
	{"Page of Swords", 11, "Minor", "Swords", "", "", "Air", "Page"},
	{"Knight of Swords", 12, "Minor", "Swords", "", "", "Air", "Knight"},
	{"Queen of Swords", 13, "Minor", "Swords", "", "", "Air", "Queen"},
	{"King of Swords", 14, "Minor", "Swords", "", "", "Air", "King"},

	// Wands
	{"Ace of Wands", 1, "Minor", "Wands", "", "", "Fire", ""},
	{"Two of Wands", 2, "Minor", "Wands", "Aries", "Mars", "Fire", ""},
	{"Three of Wands", 3, "Minor", "Wands", "Aries", "Sun", "Fire", ""},
	{"Four of Wands", 4, "Minor", "Wands", "Aries", "Venus", "Fire", ""},
	{"Five of Wands", 5, "Minor", "Wands", "Leo", "Saturn", "Fire", ""},
	{"Six of Wands", 6, "Minor", "Wands", "Leo", "Jupiter", "Fire", ""},
	{"Seven of Wands", 7, "Minor", "Wands", "Leo", "Mars", "Fire", ""},
	{"Eight of Wands", 8, "Minor", "Wands", "Sagittarius", "Mercury", "Fire", ""},
	{"Nine of Wands", 9, "Minor", "Wands", "Sagittarius", "Moon", "Fire", ""},
	{"Ten of Wands", 10, "Minor", "Wands", "Sagittarius", "Saturn", "Fire", ""},
	{"Page of Wands", 11, "Minor", "Wands", "", "", "Fire", "Page"},
	{"Knight of Wands", 12, "Minor", "Wands", "", "", "Fire", "Knight"},
	{"Queen of Wands", 13, "Minor", "Wands", "", "", "Fire", "Queen"},
	{"King of Wands", 14, "Minor", "Wands", "", "", "Fire", "King"},
}

func DrawCards(num int) (drawnCards []TarotCard) {
	tarotDeck := append([]TarotCard(nil), Deck...) // copy the deck
	drawnCards = []TarotCard{}

	for i := 0; i < num; i++ {
		if len(tarotDeck) == 0 {
			break // prevent panic if num > deck size
		}

		randomIndex := rand.Intn(len(tarotDeck))
		drawn := tarotDeck[randomIndex]

		drawnCards = append(drawnCards, drawn)
		tarotDeck = append(tarotDeck[:randomIndex], tarotDeck[randomIndex+1:]...) // remove the card
	}

	return drawnCards
}
