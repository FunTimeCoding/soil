package constant

import "regexp"

const (
	Path        = "/proc/mdstat"
	ReadCommand = "cat /proc/mdstat"
)

const FailedMarker = "(F)"

var (
	Header  = regexp.MustCompile(`^(md\S+) : (\S+) (?:\(\S+\) )?(\S+) (.+)$`)
	Members = regexp.MustCompile(`\[(\d+)/(\d+)\]`)
	Rebuild = regexp.MustCompile(`\b(recovery|resync|reshape) =`)
)
