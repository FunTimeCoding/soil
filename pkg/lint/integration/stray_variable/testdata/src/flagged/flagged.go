package flagged

import (
	"errors"
	"regexp"
)

var (
	Sentinel = errors.New("boom")
	pattern  = regexp.MustCompile(`\d+`)
	Table    = []string{"one", "two"}
	Version  = "v1"
)

func Exercise() (error, *regexp.Regexp, []string, string) {
	return Sentinel, pattern, Table, Version
}
