package constant

import "regexp"

var EscapePattern = regexp.MustCompile(
	`\x1b(?:\[[0-9;?]*[a-zA-Z]|\][^\x07]*\x07)`,
)

var TracePrefixPattern = regexp.MustCompile(
	`^(\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}\.\d+Z) (\d\d[OE]\+?) ?`,
)
