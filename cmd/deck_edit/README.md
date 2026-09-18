# deck_edit

The deck editor, outside the adventure.

The original ships a standalone deck builder — `Deck.exe`, with `Menus.txt`'s
`@DECKSURFACE_STANDALONE` surface — and the manual says its Inventory holds
"every card you can put into a deck — every Magic: The Gathering card included
in the game!" (p.139). This runs the same screen a city opens, against the whole
card database and with no city behind it: no prices, no selling, nothing to
unlock first.

```bash
go run ./cmd/deck_edit             # 4 copies of every card
go run ./cmd/deck_edit -copies 12  # more copies, for building a full deck
```

Useful for checking anything card-dependent without playing to the card first:
the filters, the deck rules, the screen itself.
