package serializers

import (
	"solitaire/game/board"
	"solitaire/game/deck"
	"solitaire/game/stacks"
	"testing"
)

func TestSerializeDeck(t *testing.T) {
	// Test with an empty deck
	emptyDeck := deck.Cards{}
	emptyDeckResponse := SerializeDeck(emptyDeck, 0)

	if emptyDeckResponse.CurrentCard != nil {
		t.Error("Expected CurrentCard to be nil for empty deck")
	}
	if emptyDeckResponse.CardsFlipped != 52 {
		t.Error("Expected CardsFlipped to be 52 for empty deck")
	}
	if emptyDeckResponse.CardsRemaining != 0 {
		t.Error("Expected CardsRemaining to be 0 for empty deck")
	}

	// Test with a full deck
	fullDeck := deck.NewDeck()
	fullDeckResponse := SerializeDeck(fullDeck, 0)

	if fullDeckResponse.CurrentCard.Value != fullDeck[0].Value || fullDeckResponse.CurrentCard.Suit != fullDeck[0].Suit {
		t.Error("CurrentCard Value or Suit does not match expected values")
	}
	if fullDeckResponse.CardsFlipped != 0 {
		t.Error("Expected CardsFlipped to be 0 for full deck")
	}
	if fullDeckResponse.CardsRemaining != 52 {
		t.Error("Expected CardsRemaining to be 52 for full deck")
	}

	// Test with a partially flipped deck
	partialDeck := deck.NewDeck()
	partialDeck = partialDeck[10:] // Simulate 10 cards flipped
	partialDeckResponse := SerializeDeck(partialDeck, 0)

	if partialDeckResponse.CardsFlipped != 10 {
		t.Error("Expected CardsFlipped to be 10")
	}
	if partialDeckResponse.CardsRemaining != 42 {
		t.Error("Expected CardsRemaining to be 42")
	}
}

func TestSerializeStacks(t *testing.T) {
	// Test with empty stacks
	emptyStacks := stacks.NewStacks()
	emptyStacksResponse := SerializeStacks(emptyStacks)

	if len(emptyStacksResponse) != 4 {
		t.Error("Expected 4 stacks in response")
	}
	for _, stackResponse := range emptyStacksResponse {
		if stackResponse.TopCard != nil {
			t.Error("Expected TopCard to be nil for empty stacks")
		}
	}

	// Test with stacks containing cards
	stacksWithCards := stacks.NewStacks()
	card1, _ := deck.NewCard(1, "Hearts", true)
	card2, _ := deck.NewCard(2, "Spades", true)
	stacksWithCards[0] = append(stacksWithCards[0], card1) // Add card to the first stack
	stacksWithCards[1] = append(stacksWithCards[1], card2) // Add card to the second stack
	stacksWithCardsResponse := SerializeStacks(stacksWithCards)

	// Check the first stack
	if stacksWithCardsResponse[0].TopCard.Value != 1 || stacksWithCardsResponse[0].TopCard.Suit != "Hearts" {
		t.Error("TopCard for the first stack does not match expected values")
	}
	// Check the second stack
	if stacksWithCardsResponse[1].TopCard.Value != 2 || stacksWithCardsResponse[1].TopCard.Suit != "Spades" {
		t.Error("TopCard for the second stack does not match expected values")
	}
	// Check if the other stacks are empty
	if stacksWithCardsResponse[2].TopCard != nil || stacksWithCardsResponse[3].TopCard != nil {
		t.Error("Expected TopCard to be nil for empty stacks")
	}
}

func TestSerializeBoard(t *testing.T) {
	// Test with an empty board
	emptyBoard := board.NewBoard()
	emptyBoardResponse := SerializeBoard(emptyBoard)
	if len(emptyBoardResponse.Columns) != board.NumColumns {
		t.Error("Incorrect number of columns in serialized board")
	}

	for _, column := range emptyBoardResponse.Columns {
		if column.HiddenCards != 0 {
			t.Error("Expected 0 hidden cards in empty column")
		}
		if len(column.VisibleCards) != 0 {
			t.Error("Expected 0 visible cards in empty column")
		}
	}

	// Test with a board containing cards
	boardWithCards := board.NewBoard()
	card1, _ := deck.NewCard(10, "Hearts", true)
	card2, _ := deck.NewCard(11, "Spades", false)
	card3, _ := deck.NewCard(12, "Diamonds", true)
	boardWithCards[0] = append(boardWithCards[0], card1, card2) // Add cards to the first column
	boardWithCards[1] = append(boardWithCards[1], card3)        // Add a card to the second column
	boardWithCardsResponse := SerializeBoard(boardWithCards)

	// Check the first column
	if boardWithCardsResponse.Columns[0].HiddenCards != 1 {
		t.Errorf("Expected 1 hidden card in the first column, got %d", boardWithCardsResponse.Columns[0].HiddenCards)
	}
	if len(boardWithCardsResponse.Columns[0].VisibleCards) != 1 {
		t.Errorf("Expected 1 visible card in the first column, got %d", len(boardWithCardsResponse.Columns[0].VisibleCards))
	}
	if boardWithCardsResponse.Columns[0].VisibleCards[0].Value != 10 ||
		boardWithCardsResponse.Columns[0].VisibleCards[0].Suit != "Hearts" {
		t.Error("Visible card in the first column does not match expected values")
	}

	// Check the second column
	if boardWithCardsResponse.Columns[1].HiddenCards != 0 {
		t.Errorf("Expected 0 hidden cards in the second column, got %d", boardWithCardsResponse.Columns[1].HiddenCards)
	}
	if len(boardWithCardsResponse.Columns[1].VisibleCards) != 1 {
		t.Errorf("Expected 1 visible card in the second column, got %d", len(boardWithCardsResponse.Columns[1].VisibleCards))
	}
	if boardWithCardsResponse.Columns[1].VisibleCards[0].Value != 12 ||
		boardWithCardsResponse.Columns[1].VisibleCards[0].Suit != "Diamonds" {
		t.Error("Visible card in the second column does not match expected values")
	}

	// ... (Add more checks for other columns and scenarios)
}

// func TestSerializeBoard(t *testing.T) {
// Making half the board hidden and making sure that the board is searlized correctly.
// 	d := deck.NewDeck()
// 	b := board.NewBoard()

// 	for i := 0; i < 7; i++ {
// 		for j := 0; j < 7; j++ {
// 			b[i] = append(b[i], d[i+(j*7)])
// 			if j > i {
// 				b[i][j].Shown = true
// 			}
// 		}
// 	}

// 	SerializeGame()
// 	columnCards, numHidden := b.GetUserResponse()

// 	for i, column := range columnCards {
// 		fmt.Println(column)
// 		fmt.Println(numHidden[i])
// 	}
// 	t.Error("in progress")
// }
