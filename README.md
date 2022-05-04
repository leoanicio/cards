## This repo contains a public API to handle deck of cards using Go. 

## Installation

- To install Go in your system, please follow the steps in https://go.dev/doc/install

- After installed, download or clone this repo to your computer

- To install the necessary packages, run `go mod download` from inside the project folder

## Execution

To run the api execute the following from inside the project folder:
- `go run .\cmd\backend\`  

To test the api, we provide three endpoints:

- /create
  - Used to create a new deck. The user can opt to create a standard deck of 52 cards or a custom deck from a given set of cards.
  - Method: POST
  - Body:
    - shuffled: bool -> If the deck is to be shuffled or not
    - cards: []string -> The cards of the deck. If empty, a standard deck of 52 cards will be created.


- /get/:deck_id
  - Used to return the deck and its cards to the user.
  - Method: GET
  - Parameters: 
    - deck_id: string -> The id of the deck to open

- /draw
  - Used to draw a card from a deck
  - Method: POST
  - Body:
    - deck_id: string -> The id of the deck to draw
    - ammount: int -> The number of cards to draw

## Tests

To run the unit tests execute the following from inside the project folder:
- `go test .\cmd\backend\`  