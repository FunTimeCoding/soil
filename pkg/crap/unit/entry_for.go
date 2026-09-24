package unit

import (
	"github.com/funtimecoding/soil/pkg/crap/entry"
	"github.com/funtimecoding/soil/pkg/crap/function"
)

func entryFor(
	packagePath string,
	file string,
	name string,
	complexity int,
	coverage float64,
) *entry.Entry {
	return entry.New(
		function.New(packagePath, file, 1, name, complexity),
		coverage,
	)
}
