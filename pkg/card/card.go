package card

type Card struct {
	Value    string `json:"value"`
	Suit     string `json:"suit"`
	Code     string `json:"code"`
	position int
}

// NewCard creates and returns a new instance of the Card object
func NewCard(code, suit, value string, pos int) Card {
	card := Card{
		Value:    value,
		Suit:     suit,
		Code:     code,
		position: pos,
	}

	return card
}
