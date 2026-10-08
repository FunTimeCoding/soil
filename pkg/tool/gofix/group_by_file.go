package gofix

import (
	"github.com/funtimecoding/soil/pkg/system"
	"go/token"
	"strings"
)

func groupByFile(
	fileSet *token.FileSet,
	edits []Edit,
	directory string,
) map[string][]FileEdit {
	result := make(map[string][]FileEdit)
	workingDirectory := directory

	if workingDirectory == "" {
		workingDirectory = system.WorkDirectory()
	}

	for _, e := range edits {
		position := fileSet.Position(e.position)
		endPosition := fileSet.Position(e.end)
		path := position.Filename

		if !strings.HasPrefix(path, workingDirectory) {
			continue
		}

		result[path] = append(
			result[path],
			FileEdit{
				offset:  position.Offset,
				length:  endPosition.Offset - position.Offset,
				newText: e.newText,
			},
		)
	}

	return result
}
