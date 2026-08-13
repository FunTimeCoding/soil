package gofix

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/lint/output"
)

func RunFormatFixWithDirectory(
	patterns []string,
	directory string,
	r *output.Results,
) {
	if len(patterns) == 0 {
		patterns = []string{"./..."}
	}

	all, _ := Load(directory, patterns)
	changed := findFormatEdits(all, r, true)

	if len(changed) == 0 {
		return
	}

	writeDestinationFiles(changed)

	for pass := range 5 {
		all, _ = Load(directory, patterns)
		changed = findFormatEdits(all, r, false)

		if len(changed) == 0 {
			return
		}

		writeDestinationFiles(changed)

		for path := range changed {
			errors.Printf("pass %d: still changing %s\n", pass+2, path)
		}
	}

	errors.Printf("format fix did not converge within 6 passes\n")
}
