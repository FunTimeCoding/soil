package constant

import "regexp"

const (
	ReflowWidth          = 80
	FixtureReflowWidth   = 40
	Backtick             = '`'
	Whitespace           = " \t\n\r\v\f"
	HorizontalWhitespace = " \t"
	HardLineBreakShape   = "hard line break"

	ReflowKey         = "reflow"
	ReflowText        = "Line exceeds the column width"
	ReflowRefusedKey  = "reflow_refused"
	ReflowRefusedText = "Reflow refused -"
)

var BlockOpenerPattern = regexp.MustCompile(`^>|^(?:[*+-]|#{1,6}|\d+[).])$`)
