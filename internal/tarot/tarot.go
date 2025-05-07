package tarot

import (
	"fmt"
	"math/rand"
)

type TarotCard struct {
	Name      string
	Number    int
	Arcana    string
	Suite     string
	Zodiac    string
	Planet    string
	Element   string
	Rank      int
	ImagePath string
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

func init() {
	suits := map[string]string{
		"Cups":      "Water",
		"Pentacles": "Earth",
		"Swords":    "Air",
		"Wands":     "Fire",
	}

	ranks := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Page", "Knight", "Queen", "King"}

	for suit, element := range suits {
		for i, rank := range ranks {
			imagePath := fmt.Sprintf("/templates/assets/cards/%s%02d.jpg", suit, i+1)
			cardName := fmt.Sprintf("%s of %s", rank, suit)

			Deck = append(Deck, TarotCard{
				Name:      cardName,
				Number:    i + 1,
				Arcana:    "Minor",
				Suite:     suit,
				Zodiac:    "",
				Planet:    "",
				Element:   element,
				Rank:      i + 1,
				ImagePath: imagePath,
			})
		}
	}
}

var Deck = []TarotCard{
	// Major Arcana
	{"The Fool", 0, "Major", "", "", "Uranus", "Air", 0, "/templates/assets/cards/00-TheFool.jpg"},
	{"The Magician", 1, "Major", "", "", "Mercury", "Air", 0, "/templates/assets/cards/01-TheMagician.jpg"},
	{"The High Priestess", 2, "Major", "", "", "Moon", "Water", 0, "/templates/assets/cards/02-TheHighPriestess.jpg"},
	{"The Empress", 3, "Major", "", "", "Venus", "Earth", 0, "/templates/assets/cards/03-TheEmpress.jpg"},
	{"The Emperor", 4, "Major", "", "Aries", "", "Fire", 0, "/templates/assets/cards/04-TheEmperor.jpg"},
	{"The Hierophant", 5, "Major", "", "Taurus", "", "Earth", 0, "/templates/assets/cards/05-TheHierophant.jpg"},
	{"The Lovers", 6, "Major", "", "Gemini", "", "Air", 0, "/templates/assets/cards/06-TheLovers.jpg"},
	{"The Chariot", 7, "Major", "", "Cancer", "", "Water", 0, "/templates/assets/cards/07-TheChariot.jpg"},
	{"Strength", 8, "Major", "", "Leo", "", "Fire", 0, "/templates/assets/cards/08-Strength.jpg"},
	{"The Hermit", 9, "Major", "", "Virgo", "", "Earth", 0, "/templates/assets/cards/09-TheHermit.jpg"},
	{"Wheel of Fortune", 10, "Major", "", "", "Jupiter", "Fire", 0, "/templates/assets/cards/10-WheelOfFortune.jpg"},
	{"Justice", 11, "Major", "", "Libra", "", "Air", 0, "/templates/assets/cards/11-Justice.jpg"},
	{"The Hanged Man", 12, "Major", "", "", "Neptune", "Water", 0, "/templates/assets/cards/12-TheHangedMan.jpg"},
	{"Death", 13, "Major", "", "Scorpio", "", "Water", 0, "/templates/assets/cards/13-Death.jpg"},
	{"Temperance", 14, "Major", "", "Sagittarius", "", "Fire", 0, "/templates/assets/cards/14-Temperance.jpg"},
	{"The Devil", 15, "Major", "", "Capricorn", "", "Earth", 0, "/templates/assets/cards/15-TheDevil.jpg"},
	{"The Tower", 16, "Major", "", "", "Mars", "Fire", 0, "/templates/assets/cards/16-TheTower.jpg"},
	{"The Star", 17, "Major", "", "Aquarius", "", "Air", 0, "/templates/assets/cards/17-TheStar.jpg"},
	{"The Moon", 18, "Major", "", "Pisces", "", "Water", 0, "/templates/assets/cards/18-TheMoon.jpg"},
	{"The Sun", 19, "Major", "", "", "Sun", "Fire", 0, "/templates/assets/cards/19-TheSun.jpg"},
	{"Judgement", 20, "Major", "", "", "Pluto", "Fire", 0, "/templates/assets/cards/20-Judgement.jpg"},
	{"The World", 21, "Major", "", "", "Saturn", "Earth", 0, "/templates/assets/cards/21-TheWorld.jpg"},
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
