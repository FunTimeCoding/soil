package constant

import "github.com/funtimecoding/soil/pkg/identity"

var Identity = identity.New(
	"goreplace",
	"Apply search/replace blocks from stdin to one file, all or none",
	"goreplace",
)

const (
	SearchMarker  = "<<<<<<< SEARCH"
	SpanMarker    = "<<<<<<< SPAN"
	AfterMarker   = "<<<<<<< AFTER"
	BeforeMarker  = "<<<<<<< BEFORE"
	UntilMarker   = "------- UNTIL"
	ThroughMarker = "------- THROUGH"
	ToEndMarker   = "------- TO END"
	DividerMarker = "======="
	ReplaceMarker = ">>>>>>> REPLACE"
)

var (
	Openings = []string{SearchMarker, SpanMarker, AfterMarker, BeforeMarker}
	Closings = []string{UntilMarker, ThroughMarker, ToEndMarker}
)

const HintLimit = 5
const HintFormat = "%s - %s"
const (
	NoBlocks         = "no blocks on stdin - expected <<<<<<< SEARCH, SPAN, AFTER or BEFORE, then =======, then >>>>>>> REPLACE"
	StrayText        = "line %d: expected <<<<<<< SEARCH, SPAN, AFTER or BEFORE, got %q"
	MissingDivider   = "block %d: no ======= before the end of input"
	MissingReplace   = "block %d: no >>>>>>> REPLACE before the end of input"
	MissingClosing   = "block %d: no ------- UNTIL, THROUGH or TO END before ======="
	EmptySearch      = "block %d: empty search"
	EmptyEnd         = "block %d: empty end after %s"
	TextAfterToEnd   = "block %d: nothing may follow ------- TO END before ======="
	MultilineAnchor  = "block %d: an AFTER or BEFORE anchor is one line"
	EmptyInsert      = "block %d: nothing to insert"
	EndNotFound      = "block %d: end not found after the start at line %d"
	NotFound         = "block %d: search not found"
	FirstLineAt      = "its first line occurs at line %s"
	FirstLineAbsent  = "its first line does not occur either"
	WhitespaceOnly   = "it matches when whitespace is ignored - check indentation and trailing spaces"
	Ambiguous        = "block %d: search matches %d times, at lines %s - add surrounding lines until it is unique"
	Overlap          = "blocks %d and %d overlap at line %d"
	AnchorRemoved    = "block %d: its anchor at line %d is inside block %d, which changes it"
	Failed           = "%s: nothing written\n%s"
	BlockHeader      = "@@ block %d, line %d"
	Applied          = "%s: %d blocks applied"
	DryRunApplied    = "%s: %d blocks would apply (dry run, nothing written)"
	RemovedPrefix    = "-"
	AddedPrefix      = "+"
	LineSeparator    = ", "
)
