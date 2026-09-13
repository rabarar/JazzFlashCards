Photo-extracted Guitar Chord Flashcards — v7

New: Practice fret range
------------------------
- Added Min / Max fret controls (1–18).
- Default remains 1–18, preserving prior behavior.
- Random chord + shape now selects only root/shape placements whose entire
  voicing fits inside the selected range.
- If a voicing has multiple octave placements that fit, Random may use any
  valid placement.
- The half-step Up / Down buttons now stop at the selected fret-range limits.
- Prev/Next shape and clicking a shape attempt to place that voicing inside
  the current range automatically.
- If the selected range is too narrow for the current voicing, the UI reports
  that condition and Random skips impossible placements.
- Fret-range preferences persist in localStorage across reloads.

The chord-shape JSON data is unchanged from v6.
