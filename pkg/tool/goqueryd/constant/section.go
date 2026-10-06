package constant

import "regexp"

const (
	BlockParagraph   = "paragraph"
	BlockTable       = "table"
	BlockList        = "list"
	BlockCode        = "code"
	BlockHeading     = "heading"
	BlockFrontMatter = "front matter"
	FrontMatterFence = "---"
	HeadingMark      = "#"
	MarkupCharacters = "*_`>"

	DeepestHeadingLevel = 6
	ProbeHeading        = " Section\n"
)

var ListItemPattern = regexp.MustCompile(`^([-*+]|\d+\.)\s`)
