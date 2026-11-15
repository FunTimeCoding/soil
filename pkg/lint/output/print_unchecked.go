package output

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/lint/types/unchecked"
)

func PrintUnchecked(entries []*unchecked.Unchecked) {
	for _, line := range UncheckedLines(entries) {
		console.Line(line)
	}
}
