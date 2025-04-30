package tarot

import (
	"fmt"
	"reflect"
	"strings"
)

type Stats struct {
	Major     int
	Minor     int
	Pentacles int
	Swords    int
	Wands     int
	Cups      int
	Aces      int
	Twos      int
	Threes    int
	Fours     int
	Fives     int
	Sixes     int
	Sevens    int
	Eights    int
	Nines     int
	Tens      int
	Pages     int
	Knights   int
	Queens    int
	Kings     int
}

func AnalyzeSpreadTarot(s []SpreadCard, stats *Stats) {
	for _, spread := range s {
		card := spread.Card

		if card.Arcana == "Major" {
			stats.Major++
		} else if card.Arcana == "Minor" {
			stats.Minor++

			switch card.Number {
			case 1:
				stats.Aces++
			case 2:
				stats.Twos++
			case 3:
				stats.Threes++
			case 4:
				stats.Fours++
			case 5:
				stats.Fives++
			case 6:
				stats.Sixes++
			case 7:
				stats.Sevens++
			case 8:
				stats.Eights++
			case 9:
				stats.Nines++
			case 10:
				stats.Tens++
			case 11:
				stats.Pages++
			case 12:
				stats.Knights++
			case 13:
				stats.Queens++
			case 14:
				stats.Kings++
			}

			switch card.Suite {
			case "Pentacles":
				stats.Pentacles++
			case "Wands":
				stats.Wands++
			case "Cups":
				stats.Cups++
			case "Swords":
				stats.Swords++
			default:
				fmt.Println("Unknown suite for card:", card.Name)
			}
		}
	}
}

// console print the stats.
func (s Stats) Print() {
	fmt.Printf("-------Statistics-------\n")
	fmt.Printf("Major Arcana Cards: %d\nMinor Arcana Cards: %d\n", s.Major, s.Minor)
	fmt.Printf("Cups: %d\nPentacles: %d\nSwords: %d\nWands: %d\n", s.Cups, s.Pentacles, s.Swords, s.Wands)

	// Get the type of the Stats struct
	val := reflect.ValueOf(s)
	typeOfStats := val.Type()

	// Iterate through the struct fields to find those from Aces to Kings
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldName := typeOfStats.Field(i).Name

		// Only print if field is from Aces to Kings and has a value greater than 0
		switch fieldName {
		case "Aces", "Twos", "Threes", "Fours", "Fives", "Sixes", "Sevens", "Eights", "Nines", "Tens", "Pages", "Knights", "Queens", "Kings":
			if field.Int() > 0 {
				fmt.Printf("%s: %d\n", fieldName, field.Int())
			}
		}
	}
}

// create string for use in a text file.
func (s Stats) String() string {
	var sb strings.Builder

	sb.WriteString("\nStatistics: \n-------------\n")
	sb.WriteString(fmt.Sprintf("Major Arcana Cards: %d\nMinor Arcana Cards: %d\n", s.Major, s.Minor))
	sb.WriteString(fmt.Sprintf("Cups: %d\nPentacles: %d\nSwords: %d\nWands: %d\n", s.Cups, s.Pentacles, s.Swords, s.Wands))

	// Get the type of the Stats struct
	val := reflect.ValueOf(s)
	typeOfStats := val.Type()

	// Iterate through the struct fields to find those from Aces to Kings
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldName := typeOfStats.Field(i).Name

		// Only add to string if field is from Aces to Kings and has a value greater than 0
		switch fieldName {
		case "Aces", "Twos", "Threes", "Fours", "Fives", "Sixes", "Sevens", "Eights", "Nines", "Tens", "Pages", "Knights", "Queens", "Kings":
			if field.Int() > 0 {
				sb.WriteString(fmt.Sprintf("%s: %d\n", fieldName, field.Int()))
			}
		}
	}

	return sb.String()
}
