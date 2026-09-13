Photo-extracted Guitar Chord Flashcards — v8 (iOS Safari hardening)

Changes:
- localStorage access is now fully guarded. Safari can deny storage in private,
  file://, or restricted browsing contexts; that no longer prevents the app
  from initializing or responding.
- Removed Array.flatMap from startup code for wider Safari compatibility.
- Removed nullish-coalescing/object-spread usage from core rendering paths.
- JSON import now uses FileReader, which is more reliable across iOS Safari.
- JSON export appends the temporary link to the DOM before triggering it.
- Mobile layout was redesigned for narrow iPhone screens:
    * header stacks vertically
    * controls use 44px+ touch targets
    * fretboard scales to viewport width
    * fretboard navigation becomes a 2x2 grid
    * Note/Finger/Degree selector stacks vertically
    * safe-area insets are respected
- Added an on-page runtime warning if Safari encounters a script error instead
  of failing silently.

The chord data and practice logic are unchanged from v7.
