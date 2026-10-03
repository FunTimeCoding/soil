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

	InterruptingListKey  = "interrupting_list"
	InterruptingListText = "List starts directly under a paragraph line - join a prose dash back, or put a blank line or a colon before a real list"
	EmphasisMarkers      = "*_"
)

var BlockOpenerPattern = regexp.MustCompile(`^>|^(?:[*+-]|#{1,6}|\d+[).])$`)
