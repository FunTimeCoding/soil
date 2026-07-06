# Fixture Pattern

Test fixtures (static files used by tests) live at `<repo-root>/fixture/<subdir>/` and are
loaded via helpers in `pkg/assert/fixture`. The root is located by walking up from the test
work directory to the nearest git root - so paths are always absolute and test-location
independent.

## Helpers

```go
import "github.com/funtimecoding/soil/pkg/assert/fixture"

fixture.Path("hypertext", "test.html")   // → absolute path string
fixture.File("hypertext", "test.html")   // → *os.File (already opened)
fixture.Read("markdown", "1.md")         // → file contents as string
```

Use `Path` when passing to a function that takes a file path.
Use `File` when passing to a function that takes `*os.File` (e.g. `goquery.NewDocumentFromReader`).
Use `Read` when you need the raw string content in the test itself.

## Directory Layout

```
<repo-root>/
└── fixture/
    ├── hypertext/      # HTML test files
    ├── markdown/       # Markdown test files
    └── <subdir>/       # One subdir per domain/package under test
```

Each subdirectory groups fixtures for a single domain. Subdirectory names are string constants
so callers never hardcode bare strings.

## Subdirectory Constants

Constants keep fixture paths refactorable and prevent bare string literals in tests.

**go-library** - constants live in `pkg/system/constant/constant.go`:

```go
HypertextPath = "hypertext"
MarkdownPath  = "markdown"
```

In other repos, define equivalent constants in a dedicated directory-constant package and add
one entry per fixture subdirectory whenever a new one is created.

## Test Pattern

```go
import (
    "github.com/funtimecoding/soil/pkg/assert"
    "github.com/funtimecoding/soil/pkg/assert/fixture"
    "github.com/funtimecoding/soil/pkg/system/constant"
    "testing"
)

func TestTitle(t *testing.T) {
    assert.String(
        t,
        "Test Title",
        Title(Document(fixture.File(constant.HypertextPath, "test.html"))),
    )
}

func TestParse(t *testing.T) {
    assert.String(
        t,
        "## Example\n\n...",
        fixture.Read(constant.MarkdownPath, "1.md"),
    )
}
```

## Do Not Use `testdata/`

Do not put fixtures inside packages as `testdata/` directories. All fixtures belong at the
repo root under `fixture/`. This keeps fixtures discoverable, shareable across packages, and
consistent with the repo-root resolution pattern.

**Exception:** analyzers under `pkg/lint/analyzer/` use `testdata/` because Go's
`analysistest` framework requires that exact layout. That is the only sanctioned use.
