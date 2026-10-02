package constant

import (
	"github.com/funtimecoding/soil/pkg/identity"
	"regexp"
)

var Identity = identity.New(
	"goflow",
	"Markdown prose reflow",
	"goflow [--width n] [--check] <file> [file ...]",
)

const (
	Width        = 80
	Mode         = 0o644
	Delimiter    = "---"
	FixtureWidth = 40
	Pipe         = "|"
	Backtick     = '`'
	Whitespace   = " \t\n\r\v\f"
	Indent       = " \t"

	HardLineBreak = "hard line break"
)

var BlockOpenerPattern = regexp.MustCompile(`^>|^(?:[*+-]|#{1,6}|\d+[).])$`)
