// The deck editor, outside the adventure.
//
// The original ships one as a separate program (Deck.exe, with Menus.txt's
// @DECKSURFACE_STANDALONE surface), and its Inventory holds "every card you can
// put into a deck — every Magic: The Gathering card included in the game!"
// (manual p.139). This is the same screen a city opens, with the whole card
// database as the collection and no city behind it: nothing here is for sale,
// which is also true of the original's standalone builder — its ten menu entries
// have no Sell among them.
package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/benprew/s30/game/domain"
	"github.com/benprew/s30/game/screens"
	"github.com/benprew/s30/game/ui"
	"github.com/benprew/s30/game/ui/screenui"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth  = 1024
	screenHeight = 768
)

// fullCollection is every card in the game, which is what the original's
// standalone builder offers and what makes this screen usable for checking a
// card you do not own yet.
func fullCollection(copies int) domain.CardCollection {
	collection := domain.NewCardCollection()
	for _, card := range domain.CARDS {
		collection.AddCard(card, copies)
	}
	return collection
}

type deckEditorGame struct {
	screen        screenui.Screen
	updatePointer func()
}

func (g *deckEditorGame) Update() error {
	g.updatePointer()
	next, _, err := g.screen.Update(screenWidth, screenHeight, 1)
	if err != nil {
		return err
	}
	if next == screenui.CityScr {
		return ebiten.Termination
	}
	return nil
}

func (g *deckEditorGame) Draw(screen *ebiten.Image) {
	g.screen.Draw(screen, screenWidth, screenHeight, 1)
}

func (g *deckEditorGame) Layout(_, _ int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	copies := flag.Int("copies", 4, "copies of each card in the collection")
	flag.Parse()
	if *copies < 1 {
		log.Fatal("-copies must be at least 1")
	}

	if _, err := domain.LoadEmbeddedCardImages(); err != nil {
		log.Fatalf("Failed to load embedded card images: %v", err)
	}

	player, err := domain.NewPlayer("Deck Editor", nil, false, domain.DifficultyEasy, domain.ColorGreen)
	if err != nil {
		log.Fatalf("Failed to create player: %v", err)
	}
	player.CardCollection = fullCollection(*copies)

	// The same warm-up the game runs when a world starts. Card art is fetched and
	// cached in the background, and until it arrives the collection list draws the
	// blank frame with the card's name instead of the card.
	go domain.PreloadCardImages(domain.CollectPriorityCards(player))

	editDeckScreen, err := screens.NewEditDeckScreen(player, nil, screenWidth, screenHeight)
	if err != nil {
		log.Fatalf("Failed to create edit deck screen: %v", err)
	}

	fmt.Printf("Collection: %d cards, %d of each\n", len(domain.CARDS)**copies, *copies)

	g := &deckEditorGame{screen: editDeckScreen, updatePointer: ui.UpdatePointer}
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Deck Editor")
	if err := ebiten.RunGame(g); err != nil && err != ebiten.Termination {
		log.Fatal(err)
	}
}
