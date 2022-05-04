package api

import (
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/leoanicio/deck_handler/pkg/deck"
)

type PayloadNewDeck struct {
	Cards    *[]string `json:"cards" binding:"required"`
	Shuffled *bool     `json:"shuffled" binding:"required"`
}

type PayloadDrawCard struct {
	Deck_id *string `json:"deck_id" binding:"required"`
	Ammount *int    `json:"ammount" binding:"required"`
}

type ReponseNewDeck struct {
	Shuffled  bool   `json:"shuffled"`
	Deck_id   string `json:"deck_id"`
	Remaining int    `json:"remaining"`
}

// CreateDeck is responsible for creating a new deck
// It returns 200 and the deck information if the parameters are correct
// If the parameters are missing or invalid, it returns 500
// If also returns 500 if there is an error during the procedure
func CreateDeck(c *gin.Context) {
	payload := PayloadNewDeck{}

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(
			500,
			gin.H{"error": err.Error()},
		)
		return
	}

	fmt.Println("GOT: ", payload, c.Params)
	createdDeck, err := deck.NewDeck(*payload.Cards, *payload.Shuffled)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, ReponseNewDeck{
		Shuffled:  createdDeck.Shuffled,
		Remaining: createdDeck.Remaining,
		Deck_id:   createdDeck.Deck_id,
	})
}

// GetDeck is responsible for getting a deck from a given deck_id
// It returns 200 and the deck, if it exists
// If the parameters are missing or invalid it returns 500
// If also returns 500 if there is an error during the procedure
func GetDeck(c *gin.Context) {
	deck_id, ok := c.Params.Get("deck_id")
	if !ok {
		c.JSON(500, gin.H{"error": "Please provide the deck id"})
		return
	}

	deck, err := deck.GetDeck(deck_id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, deck)
}

// DrawCard is responsible for drawing a card from a deck
// It returns 200 and the card, if it exists
// If the parameters are missing or invalid it returns 500
// If also returns 500 if there is an error during the procedure
func DrawCard(c *gin.Context) {
	payload := PayloadDrawCard{}

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(
			500,
			gin.H{"error": err.Error()},
		)
		return
	}

	if *payload.Ammount <= 0 {
		c.JSON(500, gin.H{"error": errors.New("the number of cards to draw should be greater than 0")})
	}

	deck, err := deck.GetDeck(*payload.Deck_id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	cards, err := deck.Draw(*payload.Ammount)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, cards)
}
