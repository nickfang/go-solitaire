package serializers

import (
	"encoding/json"
	"fmt"

	"solitaire/game"
	"solitaire/game/board"
	"solitaire/game/deck"
	"solitaire/game/stacks"
)

// DeckResponse represents the serialized data for the deck
type DeckResponse struct {
	CurrentCard    *CardResponse `json:"currentCard"`
	CardsFlipped   int           `json:"cardsFlipped"`
	CardsRemaining int           `json:"cardsRemaining"`
}

// CardResponse represents the serialized data for a card
type CardResponse struct {
	Value int    `json:"value"`
	Suit  string `json:"suit"`
}

// StackResponse represents the serialized data for a stack
type StackResponse struct {
	TopCard *CardResponse `json:"topCard"`
}

// ColumnResponse represents the serialized data for a column on the board
type ColumnResponse struct {
	HiddenCards  int            `json:"hiddenCards"`
	VisibleCards []CardResponse `json:"visibleCards"`
}

// BoardResponse represents the serialized data for the board
type BoardResponse struct {
	Columns []ColumnResponse `json:"columns"`
}

type GameResponse struct {
	Deck   DeckResponse
	Stacks []StackResponse
	Board  BoardResponse
}

// SerializeDeck creates a DeckResponse from a Deck
func SerializeDeck(deck deck.Cards, currentCardIndex int) *DeckResponse {
	var currentCard *CardResponse
	if len(deck) > 0 { // Check if there are cards in the deck
		currentCard = &CardResponse{
			Value: deck[currentCardIndex].Value,
			Suit:  deck[currentCardIndex].Suit,
		}
	}
	// Assuming CardsFlipped is managed in your deck logic, replace with your actual logic
	cardsFlipped := 52 - len(deck)
	return &DeckResponse{
		CurrentCard:    currentCard,
		CardsFlipped:   cardsFlipped,
		CardsRemaining: len(deck),
	}
}

// SerializeStacks creates an array of StackResponse from Stacks
func SerializeStacks(stacks stacks.Stacks) []StackResponse {
	stackResponses := make([]StackResponse, len(stacks))
	for i, stack := range stacks {
		if len(stack) > 0 {
			topCard := stack[len(stack)-1]
			stackResponses[i] = StackResponse{
				TopCard: &CardResponse{
					Value: topCard.Value,
					Suit:  topCard.Suit,
				},
			}
		} else {
			// Handle empty stack - maybe return a null topCard or an empty CardResponse
			stackResponses[i] = StackResponse{TopCard: nil}
		}
	}
	return stackResponses
}

// SerializeBoard creates a BoardResponse from a Board
func SerializeBoard(board board.Board) *BoardResponse {
	columns := make([]ColumnResponse, len(board))
	for i, col := range board {
		column := ColumnResponse{
			HiddenCards:  0,
			VisibleCards: []CardResponse{},
		}
		for _, card := range col {
			if card.Value != -1 {
				if card.Shown {
					column.VisibleCards = append(column.VisibleCards, CardResponse{
						Value: card.Value,
						Suit:  card.Suit,
					})
				} else {
					column.HiddenCards++
				}
			}
		}
		columns[i] = column
	}
	return &BoardResponse{Columns: columns}
}

func SerializeGame(game game.Game) *GameResponse {
	deckResponse := SerializeDeck(game.Cards, game.CurrentCardIndex)
	stackResponse := SerializeStacks(game.Stacks)
	boardResponse := SerializeBoard(game.Board)

	gameResponse := GameResponse{
		Deck:   *deckResponse,
		Stacks: stackResponse,
		Board:  *boardResponse,
	}

	return &gameResponse
}

// Example usage
func main() {
	deck := deck.NewDeck()
	stacks := stacks.NewStacks()
	board := board.NewBoard()

	deckResponse := SerializeDeck(deck, 0)
	stackResponses := SerializeStacks(stacks)
	boardResponse := SerializeBoard(board)

	// Create a response data structure to hold everything
	responseData := map[string]interface{}{
		"deck":   deckResponse,
		"stacks": stackResponses,
		"board":  boardResponse,
	}

	// Encode to JSON
	jsonData, err := json.Marshal(responseData)
	if err != nil {
		fmt.Println("Error serializing data:", err)
		return
	}

	fmt.Println(string(jsonData))
}
