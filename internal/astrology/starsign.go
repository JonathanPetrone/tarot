package astrology

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
	From    string // You can convert to time.Time if needed
	To      string
}

var StarSigns = []StarSign{
	{"Aries", Fire, "Mars", "March 21", "April 19"},
	{"Taurus", Earth, "Venus", "April 20", "May 20"},
	{"Gemini", Air, "Mercury", "May 21", "June 20"},
	{"Cancer", Water, "Moon", "June 21", "July 22"},
	{"Leo", Fire, "Sun", "July 23", "August 22"},
	{"Virgo", Earth, "Mercury", "August 23", "September 22"},
	{"Libra", Air, "Venus", "September 23", "October 22"},
	{"Scorpio", Water, "Pluto", "October 23", "November 21"},
	{"Sagittarius", Fire, "Jupiter", "November 22", "December 21"},
	{"Capricorn", Earth, "Saturn", "December 22", "January 19"},
	{"Aquarius", Air, "Uranus", "January 20", "February 18"},
	{"Pisces", Water, "Neptune", "February 19", "March 20"},
}
