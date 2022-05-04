package deck

type cardCodeMapping struct {
	name     string
	position int
	code     string
}

var valuesMapping = []cardCodeMapping{
	{
		name:     "ACE",
		position: 0,
		code:     "A",
	},
	{
		name:     "2",
		position: 1,
		code:     "2",
	},
	{
		name:     "3",
		position: 2,
		code:     "3",
	},
	{
		name:     "4",
		position: 3,
		code:     "4",
	},
	{
		name:     "5",
		position: 4,
		code:     "5",
	},
	{
		name:     "6",
		position: 5,
		code:     "6",
	},
	{
		name:     "7",
		position: 6,
		code:     "7",
	},
	{
		name:     "8",
		position: 7,
		code:     "8",
	},
	{
		name:     "9",
		position: 8,
		code:     "9",
	},
	{
		name:     "10",
		position: 9,
		code:     "10",
	},
	{
		name:     "JACK",
		position: 10,
		code:     "J",
	},
	{
		name:     "QUEEN",
		position: 11,
		code:     "Q",
	},
	{
		name:     "KING",
		position: 12,
		code:     "K",
	},
}

var suitsMapping = []cardCodeMapping{
	{
		name:     "SPADES",
		position: 0,
		code:     "S",
	},
	{
		name:     "DIAMONDS",
		position: 1,
		code:     "D",
	},
	{
		name:     "CLUBS",
		position: 2,
		code:     "C",
	},
	{
		name:     "HEARTS",
		position: 3,
		code:     "H",
	},
}
