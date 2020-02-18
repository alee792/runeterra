// Package main is an example of the Runeterra Client.
package main

import (
	"context"
	"fmt"

	"github.com/alee792/runeterra"
)

func main() {
	c := runeterra.NewClient(runeterra.Config{})

	ctx := context.Background()

	d, err := c.GetStaticDecklist(ctx)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Deck: %#v\n", c.FullDeck(d))

	cp, err := c.GetCardPositions(ctx)
	if err != nil {
		panic(err)
	}

	fmt.Printf("CardPositions: %#v\n", cp)
}
