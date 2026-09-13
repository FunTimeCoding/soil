package scan

import "regexp"

type Scanner struct {
	terms    []string
	lower    []string
	patterns []*regexp.Regexp
}
