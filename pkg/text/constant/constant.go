package constant

import soil "github.com/funtimecoding/soil/pkg/constant"

const (
	NoLimit int = -1

	DictionaryFile = "dictionary.dic"
	MinimumStem    = 3

	WikiDirectoryEnvironment = "WIKI_DIRECTORY"
)

const (
	DiffEqual  = " "
	DiffDelete = "-"
	DiffInsert = "+"
	DiffGap    = "@@"
)

var (
	DictionarySkip   = map[string]bool{".git": true, "tmp": true}
	DictionaryPrefix = map[string]bool{
		"Containerfile.": true,
		"Dockerfile.":    true,
	}
	DictionaryExtension = map[string]bool{
		".conf":      true,
		".envrc":     true,
		".gitignore": true,
		".go":        true,
		".json":      true,
		".md":        true,
		".rego":      true,
		".sh":        true,
		".txt":       true,
		".xml":       true,
		".yaml":      true,
		".yml":       true,
	}
	DictionaryNoExtension = map[string]bool{
		soil.ContainerFile: true,
		soil.DockerFile:    true,
	}
)

var (
	Inflections = map[string][]string{
		"ies":  {"y"},
		"es":   {""},
		"s":    {""},
		"ed":   {"", "e"},
		"ied":  {"y"},
		"ing":  {"", "e"},
		"er":   {"", "e"},
		"ier":  {"y"},
		"est":  {"", "e"},
		"iest": {"y"},
	}
	Doubling = map[string]bool{"ed": true, "ing": true, "er": true, "est": true}
)

const (
	SearchMarker    = "<<<<<<< SEARCH"
	SpanMarker      = "<<<<<<< SPAN"
	AfterMarker     = "<<<<<<< AFTER"
	BeforeMarker    = "<<<<<<< BEFORE"
	UntilMarker     = "------- UNTIL"
	ThroughMarker   = "------- THROUGH"
	ToEndMarker     = "------- TO END"
	DividerMarker   = "======="
	ReplaceMarker   = ">>>>>>> REPLACE"
	HintLimit       = 5
	HintFormat      = "%s - %s"
	NoBlocks        = "no blocks on stdin - expected <<<<<<< SEARCH, SPAN, AFTER or BEFORE, then =======, then >>>>>>> REPLACE"
	StrayText       = "line %d: expected <<<<<<< SEARCH, SPAN, AFTER or BEFORE, got %q"
	MissingDivider  = "block %d: no ======= before the end of input"
	MissingReplace  = "block %d: no >>>>>>> REPLACE before the end of input"
	MissingClosing  = "block %d: no ------- UNTIL, THROUGH or TO END before ======="
	EmptySearch     = "block %d: empty search"
	EmptyEnd        = "block %d: empty end after %s"
	TextAfterToEnd  = "block %d: nothing may follow ------- TO END before ======="
	MultilineAnchor = "block %d: an AFTER or BEFORE anchor is one line"
	EmptyInsert     = "block %d: nothing to insert"
	EndNotFound     = "block %d: end not found after the start at line %d"
	NotFound        = "block %d: search not found"
	FirstLineAt     = "its first line occurs at line %s"
	FirstLineAbsent = "its first line does not occur either"
	WhitespaceOnly  = "it matches when whitespace is ignored - check indentation and trailing spaces"
	Ambiguous       = "block %d: search matches %d times, at lines %s - add surrounding lines until it is unique"
	Overlap         = "blocks %d and %d overlap at line %d"
	AnchorRemoved   = "block %d: its anchor at line %d is inside block %d, which changes it"
	LineSeparator   = ", "
)

var (
	Openings = []string{SearchMarker, SpanMarker, AfterMarker, BeforeMarker}
	Closings = []string{UntilMarker, ThroughMarker, ToEndMarker}
)
