package main

import (
	"fmt"
	"solitaire/game"
	"solitaire/game/board"
	"solitaire/game/deck"
	"solitaire/game/stacks"
)

var CardNumDisplay = [13]string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}
var CardSuitsIcons = []string{"♠", "♥", "♣", "♦"}

func getCardDisplay(value int, suit string) string {
	if value == 0 {
		return ""
	}
	if value < 1 || value > 13 {
		return "Invalid Value for Card"
	}
	if suit != "Spades" && suit != "Hearts" && suit != "Clubs" && suit != "Diamonds" {
		return "Invalid Suit for Card"
	}

	displayValue := ""
	if value != 10 {
		displayValue = " "
	}
	displayValue += CardNumDisplay[value-1]

	switch suit {
	case "Spades":
		displayValue += CardSuitsIcons[0]
	case "Hearts":
		displayValue += CardSuitsIcons[1]
	case "Clubs":
		displayValue += CardSuitsIcons[2]
	case "Diamonds":
		displayValue += CardSuitsIcons[3]
	default:
		return "invalid suit: " + suit
		// do nothing, maybe throw error.
	}
	return displayValue
}

func DisplayCard(c deck.Card) {
	if c.Value == 0 {
		fmt.Print("    ")
		return
	}
	if c.Shown || c.Debug {
		fmt.Print(" " + getCardDisplay(c.Value, c.Suit))
	} else {
		fmt.Print("  * ")
	}
}

func DisplayBoard(b board.Board) {
	displayBoard := [7][19]deck.Card{}
	maxLen := 0 // add a space so the board isn't cramped with the deck.
	for i, column := range b {
		if len(column) > maxLen {
			maxLen = len(column)
		}
		copy(displayBoard[i][:], column)
	}
	for y := 0; y < maxLen; y++ {
		for x := 0; x < 7; x++ {
			DisplayCard(displayBoard[x][y])
		}
		fmt.Println()
	}
}

func DisplayStacks(s stacks.Stacks) {
	fmt.Print("     [")
	for _, stack := range s {
		numCards := len(stack)
		if numCards == 0 {
			fmt.Print("    ")
		} else {
			DisplayCard(stack[numCards-1])
		}
	}
	fmt.Println("]")
}

func DisplayCurrentCard(g game.Game) string {
	card, error := g.GetCurrentCard()
	if error != nil {

		if error.Error() == "no cards in the deck" {
			return ""
		}
		return error.Error()
	}
	return getCardDisplay(card.Value, card.Suit)
}

func DisplayCards(g game.Game) {
	fmt.Println(g.CurrentCardIndex+1, DisplayCurrentCard(g), len(g.Cards)-g.CurrentCardIndex)
}

func DisplayGame(g game.Game) {
	DisplayStacks(g.Stacks)
	DisplayBoard(g.Board)
	DisplayCards(g)
	if g.Debug {
		fmt.Println(g.GetDeckMoves())
	}
}

func incrementDigits(input string) string {
	var result []byte
	for _, char := range input {
		if char >= '0' && char <= '6' {
			digit := int(char - '0')
			digit++
			result = append(result, byte(digit+'0'))
		} else {
			result = append(result, byte(char))
		}
	}
	return string(result)
}

func DisplayHints(g game.Game) {
	hints := g.GetDeckHints()
	hints = append(hints, g.GetStackHints()...)
	hints = append(hints, g.GetBoardHints()...)
	for i, hint := range hints {
		hints[i] = incrementDigits(hint)
	}
	fmt.Println("Moves:", hints)
}
