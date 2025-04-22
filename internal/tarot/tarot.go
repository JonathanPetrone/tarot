package tarot

import (
	"math/rand"
)

type TarotCard struct {
	Name    string
	Number  int
	Arcana  string
	Suite   string
	Zodiac  string
	Planet  string
	Element string
	Rank    int // Rank is integer and can be 0 for Major Arcana
}

type Rank int

const (
	Ace Rank = iota + 1 // Starts from 1 for Ace
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Page
	Knight
	Queen
	King
)

func (r Rank) String() string {
	return [...]string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Page", "Knight", "Queen", "King"}[r-1]
}

type RankMeaning struct {
	Rank    Rank
	Meaning string
}

var RankMeanings = map[Rank]string{
	Ace:    "Beginnings, potential, inspiration",
	Two:    "Balance, duality, partnership",
	Three:  "Growth, creativity, groups",
	Four:   "Stability, foundation, structure",
	Five:   "Conflict, challenge, change",
	Six:    "Harmony, cooperation, adjustment",
	Seven:  "Assessment, reflection, faith",
	Eight:  "Movement, mastery, action",
	Nine:   "Fruition, attainment, nearing completion",
	Ten:    "Completion, culmination, legacy",
	Page:   "Curiosity, messages, learning",
	Knight: "Action, pursuit, energy",
	Queen:  "Nurturing, insight, maturity",
	King:   "Authority, leadership, control",
}

var Deck = []TarotCard{
	// Major Arcana
	{"The Fool", 0, "Major", "", "", "Uranus", "Air", 0},
	{"The Magician", 1, "Major", "", "", "Mercury", "Air", 0},
	{"The High Priestess", 2, "Major", "", "", "Moon", "Water", 0},
	{"The Empress", 3, "Major", "", "", "Venus", "Earth", 0},
	{"The Emperor", 4, "Major", "", "Aries", "", "Fire", 0},
	{"The Hierophant", 5, "Major", "", "Taurus", "", "Earth", 0},
	{"The Lovers", 6, "Major", "", "Gemini", "", "Air", 0},
	{"The Chariot", 7, "Major", "", "Cancer", "", "Water", 0},
	{"Strength", 8, "Major", "", "Leo", "", "Fire", 0},
	{"The Hermit", 9, "Major", "", "Virgo", "", "Earth", 0},
	{"Wheel of Fortune", 10, "Major", "", "", "Jupiter", "Fire", 0},
	{"Justice", 11, "Major", "", "Libra", "", "Air", 0},
	{"The Hanged Man", 12, "Major", "", "", "Neptune", "Water", 0},
	{"Death", 13, "Major", "", "Scorpio", "", "Water", 0},
	{"Temperance", 14, "Major", "", "Sagittarius", "", "Fire", 0},
	{"The Devil", 15, "Major", "", "Capricorn", "", "Earth", 0},
	{"The Tower", 16, "Major", "", "", "Mars", "Fire", 0},
	{"The Star", 17, "Major", "", "Aquarius", "", "Air", 0},
	{"The Moon", 18, "Major", "", "Pisces", "", "Water", 0},
	{"The Sun", 19, "Major", "", "", "Sun", "Fire", 0},
	{"Judgement", 20, "Major", "", "", "Pluto", "Fire", 0},
	{"The World", 21, "Major", "", "", "Saturn", "Earth", 0},

	// Minor Arcana: Cups
	{"Ace of Cups", 1, "Minor", "Cups", "", "", "Water", 1},
	{"Two of Cups", 2, "Minor", "Cups", "Cancer", "Venus", "Water", 2},
	{"Three of Cups", 3, "Minor", "Cups", "Cancer", "Mercury", "Water", 3},
	{"Four of Cups", 4, "Minor", "Cups", "Cancer", "Moon", "Water", 4},
	{"Five of Cups", 5, "Minor", "Cups", "Scorpio", "Mars", "Water", 5},
	{"Six of Cups", 6, "Minor", "Cups", "Scorpio", "Sun", "Water", 6},
	{"Seven of Cups", 7, "Minor", "Cups", "Scorpio", "Venus", "Water", 7},
	{"Eight of Cups", 8, "Minor", "Cups", "Pisces", "Saturn", "Water", 8},
	{"Nine of Cups", 9, "Minor", "Cups", "Pisces", "Jupiter", "Water", 9},
	{"Ten of Cups", 10, "Minor", "Cups", "Pisces", "Mars", "Water", 10},
	{"Page of Cups", 11, "Minor", "Cups", "", "", "Water", 11},
	{"Knight of Cups", 12, "Minor", "Cups", "", "", "Water", 12},
	{"Queen of Cups", 13, "Minor", "Cups", "", "", "Water", 13},
	{"King of Cups", 14, "Minor", "Cups", "", "", "Water", 14},

	// Minor Arcana: Pentacles
	{"Ace of Pentacles", 1, "Minor", "Pentacles", "", "", "Earth", 1},
	{"Two of Pentacles", 2, "Minor", "Pentacles", "Capricorn", "Jupiter", "Earth", 2},
	{"Three of Pentacles", 3, "Minor", "Pentacles", "Capricorn", "Mars", "Earth", 3},
	{"Four of Pentacles", 4, "Minor", "Pentacles", "Taurus", "Sun", "Earth", 4},
	{"Five of Pentacles", 5, "Minor", "Pentacles", "Taurus", "Mercury", "Earth", 5},
	{"Six of Pentacles", 6, "Minor", "Pentacles", "Taurus", "Moon", "Earth", 6},
	{"Seven of Pentacles", 7, "Minor", "Pentacles", "Virgo", "Saturn", "Earth", 7},
	{"Eight of Pentacles", 8, "Minor", "Pentacles", "Virgo", "Sun", "Earth", 8},
	{"Nine of Pentacles", 9, "Minor", "Pentacles", "Virgo", "Venus", "Earth", 9},
	{"Ten of Pentacles", 10, "Minor", "Pentacles", "Capricorn", "Mercury", "Earth", 10},
	{"Page of Pentacles", 11, "Minor", "Pentacles", "", "", "Earth", 11},
	{"Knight of Pentacles", 12, "Minor", "Pentacles", "", "", "Earth", 12},
	{"Queen of Pentacles", 13, "Minor", "Pentacles", "", "", "Earth", 13},
	{"King of Pentacles", 14, "Minor", "Pentacles", "", "", "Earth", 14},

	// Minor Arcana: Swords
	{"Ace of Swords", 1, "Minor", "Swords", "", "", "Air", 1},
	{"Two of Swords", 2, "Minor", "Swords", "Libra", "Moon", "Air", 2},
	{"Three of Swords", 3, "Minor", "Swords", "Libra", "Saturn", "Air", 3},
	{"Four of Swords", 4, "Minor", "Swords", "Libra", "Jupiter", "Air", 4},
	{"Five of Swords", 5, "Minor", "Swords", "Aquarius", "Venus", "Air", 5},
	{"Six of Swords", 6, "Minor", "Swords", "Aquarius", "Mercury", "Air", 6},
	{"Seven of Swords", 7, "Minor", "Swords", "Aquarius", "Moon", "Air", 7},
	{"Eight of Swords", 8, "Minor", "Swords", "Gemini", "Jupiter", "Air", 8},
	{"Nine of Swords", 9, "Minor", "Swords", "Gemini", "Mars", "Air", 9},
	{"Ten of Swords", 10, "Minor", "Swords", "Gemini", "Sun", "Air", 10},
	{"Page of Swords", 11, "Minor", "Swords", "", "", "Air", 11},
	{"Knight of Swords", 12, "Minor", "Swords", "", "", "Air", 12},
	{"Queen of Swords", 13, "Minor", "Swords", "", "", "Air", 13},
	{"King of Swords", 14, "Minor", "Swords", "", "", "Air", 14},

	// Minor Arcana: Wands
	{"Ace of Wands", 1, "Minor", "Wands", "", "", "Fire", 1},
	{"Two of Wands", 2, "Minor", "Wands", "Aries", "Mars", "Fire", 2},
	{"Three of Wands", 3, "Minor", "Wands", "Aries", "Sun", "Fire", 3},
	{"Four of Wands", 4, "Minor", "Wands", "Leo", "Venus", "Fire", 4},
	{"Five of Wands", 5, "Minor", "Wands", "Leo", "Saturn", "Fire", 5},
	{"Six of Wands", 6, "Minor", "Wands", "Leo", "Jupiter", "Fire", 6},
	{"Seven of Wands", 7, "Minor", "Wands", "Leo", "Mars", "Fire", 7},
	{"Eight of Wands", 8, "Minor", "Wands", "Sagittarius", "Mercury", "Fire", 8},
	{"Nine of Wands", 9, "Minor", "Wands", "Sagittarius", "Moon", "Fire", 9},
	{"Ten of Wands", 10, "Minor", "Wands", "Sagittarius", "Saturn", "Fire", 10},
	{"Page of Wands", 11, "Minor", "Wands", "", "", "Fire", 11},
	{"Knight of Wands", 12, "Minor", "Wands", "", "", "Fire", 12},
	{"Queen of Wands", 13, "Minor", "Wands", "", "", "Fire", 13},
	{"King of Wands", 14, "Minor", "Wands", "", "", "Fire", 14},
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
