package astrology

import (
	"fmt"
	"time"
)

type Element int

const (
	Fire Element = iota
	Earth
	Air
	Water
)

func (e Element) String() string {
	return [...]string{"Fire", "Earth", "Air", "Water"}[e]
}

type StarSign struct {
	Name    string
	Element Element
	Planet  string
	From    time.Time // You can convert to time.Time if needed
	To      time.Time
}

var StarSigns = []StarSign{
	{"Aries", Fire, "Mars", parseDate("March 21"), parseDate("April 19")},
	{"Taurus", Earth, "Venus", parseDate("April 20"), parseDate("May 20")},
	{"Gemini", Air, "Mercury", parseDate("May 21"), parseDate("June 20")},
	{"Cancer", Water, "Moon", parseDate("June 21"), parseDate("July 22")},
	{"Leo", Fire, "Sun", parseDate("July 23"), parseDate("August 22")},
	{"Virgo", Earth, "Mercury", parseDate("August 23"), parseDate("September 22")},
	{"Libra", Air, "Venus", parseDate("September 23"), parseDate("October 22")},
	{"Scorpio", Water, "Pluto", parseDate("October 23"), parseDate("November 21")},
	{"Sagittarius", Fire, "Jupiter", parseDate("November 22"), parseDate("December 21")},
	{"Capricorn", Earth, "Saturn", parseDate("December 22"), parseDate("January 19")},
	{"Aquarius", Air, "Uranus", parseDate("January 20"), parseDate("February 18")},
	{"Pisces", Water, "Neptune", parseDate("February 19"), parseDate("March 20")},
}

func parseDate(date string) time.Time {
	// Layout format that corresponds to the "Month Day" format (e.g., "March 21")
	layout := "January 2"

	// Use the current year for parsing
	year := time.Now().Year()

	// Combine the date with the current year
	fullDate := fmt.Sprintf("%s %d", date, year)

	// Parse the date string
	t, err := time.Parse(layout, fullDate)
	if err != nil {
		fmt.Println("Error parsing date:", err)
		return time.Time{} // Return zero value if there's an error
	}
	return t
}
