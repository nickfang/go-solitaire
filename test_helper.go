package main



func EqualCards(card1, card2 Card) bool {
	return card1.Suit == card2.Suit && card1.Rank == card2.Rank
}