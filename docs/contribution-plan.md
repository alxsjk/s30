# s30 contribution plan

Working plan for contributing to `benprew/s30` from the public fork `alxsjk/s30`.
The private repo (`remgc-shand-hd`) stays as the art/HD lab and as a source of work that
may already be built.

## The bench

A clean clone of the public fork at the upstream's geometry (1024x768), with `mage-go`
cloned beside it and linked through `go.work`, as CONTRIBUTING requires. It lives outside
the games folder, so contribution work and lab work never share a tree.

**Every PR branch is cut from `upstream/main`, never from main**, so nothing of ours leaks
into a diff.

## The funnel — how work is chosen, in order

1. **The open items in the tracker.** They are what someone already asked for, and offered
   work lands there most easily.
2. **Anything the private fork already solved.** Check before writing anything new — the
   start-screen work is the clearest example, and the mulligan bug is the counter-example
   where we had nothing.
3. **The original's missing features.** These need an issue first, per CONTRIBUTING, since
   they are not small fixes.

## Lane 1 — every open item, grouped by who opened it

All twelve were open when this was written. Grouped so that nothing reads as a selection.

### The maintainer's own (7) — his wishlist

| # | what | where we stand |
|---|---|---|
| #10 | native text for card name and cost | comment posted. Measured: name and cost can be crisp; rules text at 150 px is a sizing problem, not a rendering one. Pairs with #12 |
| #12 | Android UI too small to read card titles | the opposite end of #10's sizing question — solve together |
| #14 | restart the game | comment posted. Spec is the original's own text: the ESC menu, then "Ready to Quit? No. Yes.", returning to the opening menu |
| #15 | saves on Android | storage work; the private fork touched `storage_js.go` for web, not Android |
| #19 | look at an opponent's hand | the capability already exists behind `-show-opponent-hand` (`main.go:48`); the work is wiring it to the cards that need it |
| #20 | multiple artworks per basic land | art pipeline; adjacent to the HD work now parked |
| #21 | binder view for the card collection | the scrolling collection and its filters are already plumbing |

### Ours (3) — already in flight

| # | what | where we stand |
|---|---|---|
| #27 | test-flaky only tests one package, Android CI mutates go.mod | our issue. benprew: "I like both of those changes… submit a PR and I'll approve it" |
| #29 | the PR answering #27 | open and mergeable. The test workflow sits at `action_required` — it needs his approval to run, because it is a first-time contributor's PR |
| #31 | start-screen work | our issue. This is the Lane 3 model — see below |

### Other people's (2)

| # | what | where we stand |
|---|---|---|
| #17 | LLM policy for CONTRIBUTING.md | opened by misha-cilantro. benprew answered in August: "feel free to use AI/LLMs in your PRs, I try to review PRs on their merits not how they were constructed" — with a caveat about low-quality volume. He said he would write the policy and has not. Offering a draft is cheap and welcome |
| #30 | mulligan fix *(PR by lephlaux)* | verified against our fork: we have the identical bug and nothing to add. His fix is correct, and the AI path is unaffected because it caps at two mulligans |

## Lane 2 — the original's missing features

From the inventory built against the manual and the shipped v1.3.2 build.

- **Systems that exist as data only**: world magics, mana links, mana taps.
- **Screens whose art is already in the repo**: status/inventory, city info, dungeon and
  castle clues.
- **Systems that do not exist**: fog of war on the world map, lair variety, riddles.
- **Thin**: the collection's filters, the challenge-and-ante options, the castle reward.

## Lane 3 — QoL

Work that is in neither the original nor the issue tracker: things found while doing
something else. **#31 is the model** — the original's opening menu, a scrolling save list,
save deletion, back navigation.

**Standing rule:** while working anywhere, note QoL gaps as they appear and propose them as
their own small issue or PR. They are the easiest thing a maintainer can accept, because
they need no design agreement — and they build the standing that Lane 2 needs.

Candidates so far:

- **A draft of the LLM policy** benprew promised in #17. He has already said AI-assisted
  PRs are judged on merit, so writing the paragraph he never got to is low risk and useful.
- **The 51 MB debug binary** — `__debug_bin.exe2604960943` is committed in `11a306d` and is
  in the upstream's history, and `.gitignore` does not cover it. A two-line PR stops it
  coming back; the 51 MB already in history is his call, not ours.
- **The map's fifth button** — City Info — plus the missing view semantics (Main / Plain / Info).
- **The in-game menu entries as their screens land** — City List (F3), Dungeon Clues (F4),
  Stats (F5), Wiz Stats (F6). Port each entry only when it opens something.
- **Resolution options** — the maintainer has said he wants these for mobile vs desktop.
- **A pass on the save/load flow** — what #31 begins.
