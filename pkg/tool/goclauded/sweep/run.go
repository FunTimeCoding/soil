package sweep

import (
	"github.com/funtimecoding/soil/pkg/system"
	"path/filepath"
)

func Run(harbor string) *Result {
	system.MakeDirectory(harbor)
	sources := collectSources()
	r := &Result{}

	for _, source := range sources {
		name := filepath.Base(source)
		destination := filepath.Join(harbor, name)

		switch synchronize(source, destination) {
		case actionSkip:
			r.Skipped++
		case actionCopy:
			copyFile(source, destination)
			r.Copied++
		case actionUpdate:
			copyFile(source, destination)
			r.Updated++
		}
	}

	return r
}
