# s30 contribution plan

Working plan for contributing to `benprew/s30` from the public fork `alxsjk/s30`.
The private repo (`remgc-shand-hd`) stays as the art/HD lab and as a source of work that
may already be built.

## The bench

A clean clone of the public fork at the upstream's geometry (1024x768), with `mage-go`
linked beside it through `go.work`, as CONTRIBUTING requires. **Every PR branch is cut
from `upstream/main`, never from main**, so nothing of ours leaks into a diff.

## The funnel — how work is chosen, in order

1. **The maintainer's open issues.** They are his wishlist, and offered work lands there
   most easily. Several we can already approach from something built in the private fork.
2. **Anything the private fork already solved.** Check before writing anything new — the
   start-screen work is the clearest example, and the mulligan bug is the counter-example
   where we had nothing.
3. **The original's missing features.** These need an issue first, per CONTRIBUTING, since
   they are not small fixes.

## Lane 1 — the maintainer's open issues

| # | what | where we stand |
|---|---|---|
| #10 | native text for card name and cost | comment posted. Measured: name and cost can be crisp; rules text at 150 px is a sizing problem, not a rendering one. Pairs with #12 |
| #14 | restart / in-game menu | comment posted. Spec is the original's own text: ESC menu, then "Ready to Quit? No. Yes.", returning to the opening menu |
| #19 | look at an opponent's hand | the capability already exists behind `-show-opponent-hand` (`main.go:48`); the work is wiring it to the cards that need it |
| #21 | binder view for the collection | the scrolling collection and filters are already plumbing |
| #12 | Android UI too small | the opposite end of #10's sizing question — solve together |
| #15 | saves on Android | storage work; the private fork touched `storage_js.go` for web, not Android |
| #20 | multiple artworks per basic land | art pipeline; adjacent to the HD work now parked |
| #17 | LLM policy in CONTRIBUTING | directly relevant to us — we use AI assistance heavily, and should say so in that thread |
| #27 | CI flakiness | our PR #29 waits on it |
| #30 | *(PR by lephlaux)* mulligan fix | verified: our fork has the identical bug and nothing of ours to offer; his fix is correct and the AI path is unaffected |

## Lane 2 — the original's missing features

From the inventory built against the manual and the shipped v1.3.2 build.

- **Systems that exist as data only**: world magics, mana links, mana taps.
- **Screens whose art is already in the repo**: status/inventory, city info, dungeon and
  castle clues.
- **Systems that do not exist**: fog of war on the world map, lair variety, riddles.
- **Thin**: the collection's filters, the challenge-and-ante options, the castle reward.

## Lane 3 — QoL

Work that is neither in the original nor in the issue tracker: things found while doing
something else. **#31 is the model** — the original's opening menu, a scrolling save list,
save deletion, back navigation.

**Standing rule:** while working anywhere, note QoL gaps as they appear and propose them as
their own small issue or PR. They are the easiest thing a maintainer can accept, because
they need no design agreement — and they build the standing needed for Lane 2.

Candidates so far:

- **The map's fifth button** — City Info — plus the missing view semantics (Main / Plain / Info).
- **The in-game menu entries as their screens land** — City List (F3), Dungeon Clues (F4),
  Stats (F5), Wiz Stats (F6). Port each entry only when it opens something.
- **Resolution options** — the maintainer has said he wants these for mobile vs desktop.
- **A pass on the save/load flow** — what #31 begins.
