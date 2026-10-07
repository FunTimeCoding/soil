package output

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/lint/types/unchecked"
)

func PrintCensus(entries []*unchecked.Unchecked) {
	for _, line := range CensusLines(entries) {
		console.Line(line)
	}
}
