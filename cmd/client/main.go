package main

import (
	"context"
	"fmt"

	"github.com/alee792/runeterra/pkg/game"
)

func main() {
	c := game.NewClient(game.Config{})

	ctx := context.Background()

	// d, err := c.GetStaticDecklist(ctx)
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Printf("Deck: %#v\n", c.FullDeck(d))

	cp, err := c.GetCardPositions(ctx)
	if err != nil {
		panic(err)
	}

	fmt.Printf("CardPositions: %#v\n", cp)
}
