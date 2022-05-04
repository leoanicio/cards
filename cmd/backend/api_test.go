package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/leoanicio/deck_handler/pkg/api"
	"github.com/leoanicio/deck_handler/pkg/card"
	"github.com/leoanicio/deck_handler/pkg/deck"
	"github.com/stretchr/testify/assert"
)

var deckIdShuffled, deckIdNotShuffled string
var router = setupRouter()

func TestCreateDeck(t *testing.T) {
	var w = httptest.NewRecorder()
	payload := api.PayloadNewDeck{
		Shuffled: &[]bool{false}[0],
		Cards:    &[]string{},
	}

	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(payload)
	if err != nil {
		log.Fatal(err)
	}

	req, err := http.NewRequest("POST", "/create", &buf)
	router.ServeHTTP(w, req)
	if err != nil {
		t.Fatal(err)
	}

	var body = api.ReponseNewDeck{}
	err = json.Unmarshal(w.Body.Bytes(), &body)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, false, body.Shuffled)
	assert.Equal(t, 52, body.Remaining)
	assert.NotNil(t, body.Deck_id)

	deckIdNotShuffled = body.Deck_id
}

func TestCreateDeckShuffled(t *testing.T) {
	var w = httptest.NewRecorder()
	payload := api.PayloadNewDeck{
		Shuffled: &[]bool{true}[0],
		Cards:    &[]string{},
	}

	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(payload)
	if err != nil {
		log.Fatal(err)
	}

	req, err := http.NewRequest("POST", "/create", &buf)
	router.ServeHTTP(w, req)
	if err != nil {
		t.Fatal(err)
	}

	var body = api.ReponseNewDeck{}
	err = json.Unmarshal(w.Body.Bytes(), &body)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, true, body.Shuffled)
	assert.Equal(t, 52, body.Remaining)
	assert.NotNil(t, body.Deck_id)

	deckIdShuffled = body.Deck_id
}

func TestCreateDeckFromCards(t *testing.T) {
	var w = httptest.NewRecorder()
	payload := api.PayloadNewDeck{
		Shuffled: &[]bool{false}[0],
		Cards:    &[]string{"KH", "KS", "AS"},
	}

	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(payload)
	if err != nil {
		log.Fatal(err)
	}

	req, err := http.NewRequest("POST", "/create", &buf)
	router.ServeHTTP(w, req)
	if err != nil {
		t.Fatal(err)
	}

	var body = api.ReponseNewDeck{}
	err = json.Unmarshal(w.Body.Bytes(), &body)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, false, body.Shuffled)
	assert.Equal(t, len(*payload.Cards), body.Remaining)
	assert.NotNil(t, body.Deck_id)
}

func TestCreateDeckFromInvalidCards(t *testing.T) {
	var w = httptest.NewRecorder()
	payload := api.PayloadNewDeck{
		Shuffled: &[]bool{false}[0],
		Cards:    &[]string{"NOTACARD"},
	}

	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(payload)
	if err != nil {
		log.Fatal(err)
	}

	req, err := http.NewRequest("POST", "/create", &buf)
	router.ServeHTTP(w, req)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 500, w.Code)
}

func TestCreateDeckMissingParameters(t *testing.T) {
	var w = httptest.NewRecorder()
	payload := map[string]*[]string{
		"Cards": &[]string{"NOTACARD"},
	}

	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(payload)
	if err != nil {
		log.Fatal(err)
	}

	req, err := http.NewRequest("POST", "/create", &buf)
	router.ServeHTTP(w, req)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 500, w.Code)
}

func TestOpenDeck(t *testing.T) {
	var w = httptest.NewRecorder()
	req, err := http.NewRequest("GET", fmt.Sprintf("/get/%s", deckIdNotShuffled), nil)
	router.ServeHTTP(w, req)
	if err != nil {
		t.Fatal(err)
	}

	var body = deck.Deck{}
	err = json.Unmarshal(w.Body.Bytes(), &body)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, 52, body.Remaining)
	assert.Equal(t, "AS", body.Cards[0].Code)
}

func TestOpenInvalidDeck(t *testing.T) {
	var w = httptest.NewRecorder()
	req, err := http.NewRequest("GET", fmt.Sprintf("/get/%s", "12345"), nil)
	router.ServeHTTP(w, req)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 500, w.Code)
}

func TestDrawCard(t *testing.T) {
	var w = httptest.NewRecorder()
	payload := api.PayloadDrawCard{
		Deck_id: &[]string{deckIdNotShuffled}[0],
		Ammount: &[]int{1}[0],
	}

	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(payload)
	if err != nil {
		log.Fatal(err)
	}

	req, err := http.NewRequest("POST", "/draw", &buf)
	router.ServeHTTP(w, req)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(w.Body)
	var body = []card.Card{}
	err = json.Unmarshal(w.Body.Bytes(), &body)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "AS", body[0].Code)
}

func TestDeckUpdatedAfterDraw(t *testing.T) {
	var w = httptest.NewRecorder()
	req, err := http.NewRequest("GET", fmt.Sprintf("/get/%s", deckIdNotShuffled), nil)
	router.ServeHTTP(w, req)
	if err != nil {
		t.Fatal(err)
	}

	var body = deck.Deck{}
	err = json.Unmarshal(w.Body.Bytes(), &body)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, 51, body.Remaining)
}

func TestDrawCardInvalidNumber(t *testing.T) {
	var w = httptest.NewRecorder()
	payload := api.PayloadDrawCard{
		Deck_id: &[]string{deckIdNotShuffled}[0],
		Ammount: &[]int{-1}[0],
	}

	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(payload)
	if err != nil {
		log.Fatal(err)
	}

	req, err := http.NewRequest("POST", "/draw", &buf)
	router.ServeHTTP(w, req)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 500, w.Code)
}

func TestDrawMoreCardsThanDeckSize(t *testing.T) {
	var w = httptest.NewRecorder()
	payload := api.PayloadDrawCard{
		Deck_id: &[]string{deckIdNotShuffled}[0],
		Ammount: &[]int{100}[0],
	}

	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(payload)
	if err != nil {
		log.Fatal(err)
	}

	req, err := http.NewRequest("POST", "/draw", &buf)
	router.ServeHTTP(w, req)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 500, w.Code)
}

func TestDrawInvalidDeck(t *testing.T) {
	var w = httptest.NewRecorder()
	payload := api.PayloadDrawCard{
		Deck_id: &[]string{"1234"}[0],
		Ammount: &[]int{1}[0],
	}

	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(payload)
	if err != nil {
		log.Fatal(err)
	}

	req, err := http.NewRequest("POST", "/draw", &buf)
	router.ServeHTTP(w, req)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 500, w.Code)
}
