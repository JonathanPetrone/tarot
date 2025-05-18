package tarot

type Spread struct {
	Name     string
	Length   int
	Template []SpreadCard
}

type SpreadCard struct {
	Position int
	Context  string
	Card     TarotCard
}

var OneCard = Spread{
	Name:   "One Card",
	Length: 1,
	Template: []SpreadCard{
		{Position: 1, Context: "The Question"},
	},
}

var PastPresentFuture = Spread{
	Name:   "Past-Present-Future",
	Length: 3,
	Template: []SpreadCard{
		{Position: 1, Context: "Past"},
		{Position: 2, Context: "Present"},
		{Position: 3, Context: "Future"},
	},
}

var CelticCross = Spread{
	Name:   "Celtic Cross",
	Length: 10,
	Template: []SpreadCard{
		{Position: 1, Context: "Current Situation"},
		{Position: 2, Context: "Challenges"},
		{Position: 3, Context: "Strengths"},
		{Position: 4, Context: "What to Focus On"},
		{Position: 5, Context: "Past Energy"},
		{Position: 6, Context: "Near Future"},
		{Position: 7, Context: "Advice"},
		{Position: 8, Context: "External Influences"},
		{Position: 9, Context: "Hopes and Fears"},
		{Position: 10, Context: "Likely Outcome"},
	},
}

func ReadSpread(s Spread) []SpreadCard {
	reading := make([]SpreadCard, s.Length)
	copy(reading, s.Template)

	drawnCards := DrawCards(s.Length)

	for i, j := range drawnCards {
		reading[i].Card = j
	}

	return reading
}
