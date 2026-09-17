package output

import "github.com/funtimecoding/soil/pkg/console"

func PrintCensus(entries []*Unchecked) {
	for _, line := range CensusLines(entries) {
		console.Line(line)
	}
}
