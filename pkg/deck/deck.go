package deck

import (
	"errors"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/google/uuid"

	"github.com/leoanicio/deck_handler/pkg/card"
)

var myDecks = make(map[string]*Deck)

type Deck struct {
	Cards     []card.Card `json:"cards"`
	Deck_id   string      `json:"deck_id"`
	Shuffled  bool        `json:"shuffled"`
	Remaining int         `json:"remaining"`
}

// GetDeck receives a deck id and check if it exists in memory
// It returns the corresponding deck, if found, otherwise it returns and error
func GetDeck(deck_id string) (*Deck, error) {
	if deck, ok := myDecks[deck_id]; ok {
		return deck, nil
	}

	return &Deck{}, errors.New("Deck does not exists")
}

// NewDeck creates and returns a new deck of cards based on the user input
// The new deck is saved in an map using the id for future use
func NewDeck(cardList []string, shuffled bool) (*Deck, error) {
	var deckCards []card.Card
	var deck_id string
	var err error
	if len(cardList) != 0 {
		deckCards, err = generateDeckFromCards(cardList, shuffled)
	} else {
		deckCards, err = generateDeckOfCards(shuffled)
	}
	if err != nil {
		return &Deck{}, err
	}

	// The id
	deck_id, err = generateDeckId()
	if err != nil {
		return &Deck{}, err
	}

	// Put it all in together
	deck := Deck{
		Cards:     deckCards,
		Shuffled:  shuffled,
		Deck_id:   deck_id,
		Remaining: len(deckCards),
	}

	myDecks[deck_id] = &deck

	return &deck, nil
}

// getDataByPosition receives a mapping and a position and returns the corresponding mapped object, if it exists
func getDataByPosition(i int, mapping []cardCodeMapping) cardCodeMapping {
	for _, metadata := range mapping {
		if metadata.position == i {
			return metadata
		}
	}
	return cardCodeMapping{}
}

// getDataByCode receives a mapping and a code and returns the corresponding mapped object, if it exists
func getDataByCode(c string, mapping []cardCodeMapping) cardCodeMapping {
	for _, metadata := range mapping {
		if metadata.code == c {
			return metadata
		}
	}
	return cardCodeMapping{}
}

// generateDeckFromCards uses the list of card codes provided by the user to create a new deck
// If the user provides an invalid car code the function returns an error
// In the end the function shuffles the deck, if requested
func generateDeckFromCards(cardList []string, shuffled bool) ([]card.Card, error) {
	cards := []card.Card{}

	for i, code := range cardList {
		cardSplit := strings.Split(code, "")
		value := cardSplit[0]
		suit := cardSplit[1]

		// Need to assert that the code exists
		valueByCode := getDataByCode(value, valuesMapping)
		suitByCode := getDataByCode(suit, suitsMapping)
		if valueByCode.code != value || suitByCode.code != suit {
			return []card.Card{}, fmt.Errorf("invalid card code: %s", code)
		}

		newCard := card.NewCard(code, suit, value, i)
		cards = append(cards, newCard)
	}

	if shuffled {
		rand.Seed(time.Now().UnixNano())

		rand.Shuffle(len(cards), func(i, j int) {
			cards[i], cards[j] = cards[j], cards[i]
		})
	}

	return cards, nil
}

// generateDeckOfCards creates and returns a standard deck of 52 cards
// It uses the position field of the suits and values mapping to keep the deck ordered 
// In the end the function shuffles the deck, if requested
func generateDeckOfCards(shuffled bool) ([]card.Card, error) {
	cards := []card.Card{}

	for i := 0; i < len(suitsMapping)*len(valuesMapping); i++ {
		var suit cardCodeMapping
		var value cardCodeMapping

		if !shuffled {
			suit = getDataByPosition(i/len(valuesMapping), suitsMapping)
			value = getDataByPosition(i%len(valuesMapping), valuesMapping)
		} else {
			suit = suitsMapping[i/len(valuesMapping)]
			value = valuesMapping[i%len(valuesMapping)]
		}

		code := fmt.Sprintf("%s%s", value.code, suit.code)
		cards = append(cards, card.NewCard(code, suit.code, value.name, i))
	}

	if shuffled {
		rand.Seed(time.Now().UnixNano())

		rand.Shuffle(len(cards), func(i, j int) {
			cards[i], cards[j] = cards[j], cards[i]
		})
	}

	return cards, nil
}

// generateDeckId creates and returns an UUID string
func generateDeckId() (string, error) {
	deck_id, err := uuid.NewRandom()

	if err != nil {
		return "", err
	}

	return deck_id.String(), nil
}

// Draw is a Deck's built in method to draw a card from the deck
// It checks if the deck is empty, or if it has the needed ammount of cards
// It returns the first card in the deck, and also updates the deck itself
func (d *Deck) Draw(n int) ([]card.Card, error) {
	if d.Remaining == 0 {
		return []card.Card{}, errors.New("this deck is empty")
	}

	if n > d.Remaining {
		return []card.Card{}, errors.New("insufficient cards in the deck")
	}

	cards := d.Cards[:n]
	d.Cards = d.Cards[n:]
	d.Remaining -= n

	return cards, nil
}
