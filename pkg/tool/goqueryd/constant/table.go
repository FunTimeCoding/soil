package constant

import "regexp"

const (
	TableRowPrefix   = "|"
	BacktickFence    = "```"
	TildeFence       = "~~~"
	TablePieceMarker = "(part %d of %d, rows %d–%d of %d)"
)

var (
	AnyHeadingPattern     = regexp.MustCompile(`^#{1,6}\s`)
	TableSeparatorPattern = regexp.MustCompile(`^\|[\s:|-]*-[\s:|-]*$`)
)
