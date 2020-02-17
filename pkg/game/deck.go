package game

import (
	"github.com/alee792/runeterra/proto"
)

// DeckCard is a card and its number of copies.
type DeckCard struct {
	proto.Card
	Copies int
}

// FullDeck contains full cards.
type FullDeck struct {
	Code  string
	Cards map[string]DeckCard
}

// FullDeck is populated from a proto.Deck.
func (c *Client) FullDeck(d *proto.Deck) *FullDeck {
	fd := &FullDeck{
		Code:  d.GetDeckCode(),
		Cards: make(map[string]DeckCard),
	}

	for k, v := range d.GetCardsInDeck() {
		fd.Cards[k] = DeckCard{
			Copies: int(v),
			Card:   c.cards[k],
		}
	}

	return fd
}
