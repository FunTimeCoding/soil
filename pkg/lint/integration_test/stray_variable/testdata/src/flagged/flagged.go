package flagged

import (
	"errors"
	"regexp"
)

var (
	Sentinel = errors.New("boom")
	pattern  = regexp.MustCompile(`\d+`)
	Table    = []string{"one", "two"}
)

func Exercise() (error, *regexp.Regexp, []string) {
	return Sentinel, pattern, Table
}
