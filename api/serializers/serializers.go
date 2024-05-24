package serializers

import (
	"encoding/json"
	"fmt"
	"solitaire/game/deck" // Assuming this is your deck package
)

// DeckResponse represents the serialized data for the deck
type DeckResponse struct {
	CurrentCard    *CardResponse `json:"currentCard"`
	CardsFlipped   int           `json:"cardsFlipped"`
	CardsRemaining int           `json:"cardsRemaining"`
}

// CardResponse represents the serialized data for a card
type CardResponse struct {
	Value string `json:"value"`
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

// SerializeDeck creates a DeckResponse from a Deck
func SerializeDeck(d *deck.Deck) *DeckResponse {
	var currentCard *CardResponse
	if len(d.Cards) > 0 { // Check if there are cards in the deck
		currentCard = &CardResponse{
			Value: d.Cards[0].Rank, // Assuming you have a Rank field in your Card struct
			Suit:  d.Cards[0].Suit,
		}
	}
	// Assuming CardsFlipped is managed in your deck logic, replace with your actual logic
	cardsFlipped := 52 - len(d.Cards)
	return &DeckResponse{
		CurrentCard:    currentCard,
		CardsFlipped:   cardsFlipped,
		CardsRemaining: len(d.Cards),
	}
}

// SerializeStacks creates an array of StackResponse from Stacks
func SerializeStacks(stacks []*deck.Card) []StackResponse {
	stackResponses := make([]StackResponse, len(stacks))
	for i, stack := range stacks {
		if stack != nil {
			stackResponses[i] = StackResponse{
				TopCard: &CardResponse{
					Value: stack.Rank,
					Suit:  stack.Suit,
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
func SerializeBoard(board [][]*deck.Card) *BoardResponse {
	columns := make([]ColumnResponse, len(board))
	for i, col := range board {
		column := ColumnResponse{
			HiddenCards:  0,
			VisibleCards: []CardResponse{},
		}
		for _, card := range col {
			if card != nil {
				if card.Shown { // Assuming Shown is a field indicating if the card is visible
					column.VisibleCards = append(column.VisibleCards, CardResponse{
						Value: card.Rank,
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

// Example usage
func main() {
	// ... (Assuming you have your deck, stacks, and board data)

	deckResponse := SerializeDeck(deck)
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
