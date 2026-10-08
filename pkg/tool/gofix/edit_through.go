package gofix

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gofix/workspace"
	"go/token"
	"sort"
)

func editThrough(
	fileSet *token.FileSet,
	edits []Edit,
	directory string,
	w *workspace.Workspace,
) {
	grouped := groupByFile(fileSet, edits, directory)

	for path, fileEdits := range grouped {
		sort.Slice(
			fileEdits,
			func(i, j int) bool {
				return fileEdits[i].offset > fileEdits[j].offset
			},
		)
		original, e := w.Read(path)

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

		if e = w.Write(path, modified); e != nil {
			errors.Printf("write %s: %s\n", path, e)
		}
	}
}
