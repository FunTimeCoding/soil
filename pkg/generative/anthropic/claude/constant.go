package claude

import "regexp"

const RecentMessageLimit = 50

var (
	markupTagPattern = regexp.MustCompile(`<[^>]+>`)
	ansiPattern      = regexp.MustCompile(`\x1b\[[0-9;]*m`)
)
