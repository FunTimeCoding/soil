# Comments

Default to no comment. Code carries its own meaning through names,
types, constants, flags, tests and call sites; a comment competes with
all of that and wins only when it says something none of them can.

The cost is not the line. It is that prose next to code reads as
authoritative, is never compiled, never tested and never re-read, and
so goes false silently while looking maintained.

## The test

**Could a reader derive this from the whole file?** Not from the line -
from the file, including the names it uses, the constants it reads, the
tests beside it and the callers it serves. If yes, the comment is
noise. Delete it.

What survives is a constraint that is true and invisible: an ordering
that matters, a unit that is not in the name, a guard whose absence
caused a real failure, a value that looks arbitrary and is derived.

## What to delete

**Restating the name.** `// New returns a new Client`. The signature
already said it.

**Narrating the next line.** `// iterate over the rows`, `// check the
error`. The code is the narration.

**Section headings inside a function.** `// --- validation ---`. If a
function needs headings it needs extracting.

**Justifying an edit.** `// switched to a pointer because the zero
value was ambiguous`. That belongs in the commit; a file is not a
changelog.

**Teaching the domain.** Paragraphs explaining the subject matter
rather than the code. Those belong in the design doc, where they can be
read whole by somebody who needs them rather than encountered by
somebody reading a function.

**Restating a constant.** `// 30 second timeout` above
`Timeout = 30 * time.Second`.

## What to graduate

A comment carrying a **design decision** - why this shape rather than
the obvious one, what was tried and rejected, how two mechanisms relate
- goes to the service's design doc, and the code
keeps nothing. Not a pointer, not a summary. The design doc is found by
the reader who needs the why; the reader of the function does not.

The tell is length. A comment that takes a paragraph to say why is
design documentation that was written in the wrong file.

## What to refactor instead

A comment that exists because the code is unclear is a **symptom**.
Removing it and leaving the code is wrong, and so is keeping it.

- explaining what a variable holds → rename the variable
- explaining what a block does → extract it and name the function
- explaining which case this handles → name the case, or make it an
  enum value with a name
- explaining a magic number → extract a named constant

Prefer the refactor. It carries the same meaning into a place the
compiler checks and the reader cannot skip.

## What earns its place

Rare, and always about a constraint rather than a description:

```go
// AutoMigrate cannot reshape an index, so the old one is dropped first.
```

```go
// Flushed before rotation, not after: the old handle keeps buffered
// writes that would land in the rotated file.
```

Both state something the file cannot show and a maintainer would
otherwise undo.

**Density matches siblings.** A package where every function carries a
paragraph and one where none do are both consistent; a file that
disagrees with its neighbours is the one to fix.

## In tests

**A comment on a test is a claim the test is not making.** Either assert
it, or split until the name carries it.

The failure is invisible because the prose reads as documentation
rather than as an unchecked assertion. A comment saying a rejected
write "leaves the cache warm" beside a test that only checks a boolean
is not describing the test - it is describing a second test that was
never written, and the reader believes it is covered.

Three dispositions, in order of preference:

- **The claim is untested** - write the assertion. The comment was a
  coverage gap wearing prose.
- **The claim is tested but unnamed** - rename the test so the name is
  the claim. `TestRejectionHoldsAtEveryBatchSize` needs nothing above it.
- **The test makes several claims** - split it. One test per claim, each
  named for its own, and the comment dissolves.

A test name is allowed to be a sentence.
`TestSpacedRetriesReopenTheCircuitSooner` and
`TestEvictionIsIndependentOfScanOrder` carry their claim in the name
and need no comment; if a claim will not fit in a name, it is more
than one claim.

What survives is the same as everywhere else: a constraint about the
*fixture* rather than the subject - why this seed value, why this
duration, what the number would otherwise be mistaken for.

## Doc comments

Exported symbols follow the same test. A doc comment that restates the
name earns nothing from being exported. Write one when the contract has
a constraint a caller must know - what is mutated, what is assumed, what
is not safe - and otherwise leave the name to do its work.
