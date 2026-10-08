package constant

import "github.com/funtimecoding/soil/pkg/identity"

var Identity = identity.New(
	"goreplace",
	"Apply search/replace blocks from stdin to one file, all or none",
	"goreplace",
)

const (
	SearchMarker  = "<<<<<<< SEARCH"
	DividerMarker = "======="
	ReplaceMarker = ">>>>>>> REPLACE"
)

const HintLimit = 5
const HintFormat = "%s - %s"
const (
	NoBlocks         = "no blocks on stdin - expected <<<<<<< SEARCH, =======, >>>>>>> REPLACE"
	StrayText        = "line %d: expected <<<<<<< SEARCH, got %q"
	MissingDivider   = "block %d: no ======= before the end of input"
	MissingReplace   = "block %d: no >>>>>>> REPLACE before the end of input"
	EmptySearch      = "block %d: empty search"
	NotFound         = "block %d: search not found"
	FirstLineAt      = "its first line occurs at line %s"
	FirstLineAbsent  = "its first line does not occur either"
	WhitespaceOnly   = "it matches when whitespace is ignored - check indentation and trailing spaces"
	Ambiguous        = "block %d: search matches %d times, at lines %s - add surrounding lines until it is unique"
	Overlap          = "blocks %d and %d overlap at line %d"
	Failed           = "%s: nothing written\n%s"
	BlockHeader      = "@@ block %d, line %d"
	Applied          = "%s: %d blocks applied"
	DryRunApplied    = "%s: %d blocks would apply (dry run, nothing written)"
	RemovedPrefix    = "-"
	AddedPrefix      = "+"
	LineSeparator    = ", "
)
