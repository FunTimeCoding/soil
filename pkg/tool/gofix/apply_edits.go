package gofix

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"go/token"
	"os"
	"sort"
)

func ApplyEdits(
	fileSet *token.FileSet,
	edits []edit,
	directory string,
	diff bool,
) {
	grouped := groupByFile(fileSet, edits, directory)

	for path, fileEdits := range grouped {
		sort.Slice(
			fileEdits,
			func(i, j int) bool {
				return fileEdits[i].offset > fileEdits[j].offset
			},
		)
		original, e := os.ReadFile(path)

		if e != nil {
			errors.Printf("read %s: %s\n", path, e)

			continue
		}

		modified := make([]byte, len(original))
		copy(modified, original)

		for _, fe := range fileEdits {
			modified = splice(
				modified,
				fe.offset,
				fe.length,
				[]byte(fe.newText),
			)
		}

		if diff {
			printDiff(path, original, modified)
		} else {
			e = os.WriteFile(path, modified, 0644)

			if e != nil {
				errors.Printf("write %s: %s\n", path, e)
			}
		}
	}
}
