# Formatting

The fix pipeline owns mechanical formatting: blank-line placement
around control structures and returns, spacing between declarations,
and call layout. Write close to the convention and let `task lint`
settle the rest - its findings name the rule they apply.

What the fixer cannot decide:

- **Function size** - aim for ~15 calls or fewer per function.
- **Line breaking** - break at 80 characters, prioritizing fewer
  total lines: break single-arg calls before multi-arg calls. When
  arguments split, each gets its own line with the closing paren
  separate.
