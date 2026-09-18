package main

import (
	"testing"

	"github.com/benprew/s30/game/domain"
	"github.com/benprew/s30/game/screens"
)

// The collection has to be the whole card set: that is the point of this tool,
// and the reason the original ships its builder as a program of its own.
func TestFullCollectionHoldsEveryCard(t *testing.T) {
	collection := fullCollection(4)
	if len(domain.CARDS) == 0 {
		t.Fatal("the card database is empty")
	}
	for _, card := range domain.CARDS {
		if got := collection.GetTotalCount(card); got != 4 {
			t.Fatalf("%s: %d copies in the collection, want 4", card.Name(), got)
		}
	}
}

// The screen has to build without a city, which is how this tool opens it: no
// prices, no selling, and nothing left that dereferences the city.
func TestDeckScreenBuildsWithoutACity(t *testing.T) {
	if _, err := domain.LoadEmbeddedCardImages(); err != nil {
		t.Fatalf("LoadEmbeddedCardImages: %v", err)
	}
	player, err := domain.NewPlayer("Deck Editor", nil, false, domain.DifficultyEasy, domain.ColorGreen)
	if err != nil {
		t.Fatalf("NewPlayer: %v", err)
	}
	player.CardCollection = fullCollection(1)

	if _, err := screens.NewEditDeckScreen(player, nil, screenWidth, screenHeight); err != nil {
		t.Fatalf("NewEditDeckScreen without a city: %v", err)
	}
}
