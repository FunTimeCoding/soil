package gofix

import (
	"github.com/funtimecoding/soil/pkg/system"
	"go/token"
	"strings"
)

func groupByFile(
	fileSet *token.FileSet,
	edits []edit,
	directory string,
) map[string][]fileEdit {
	result := make(map[string][]fileEdit)
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
			fileEdit{
				offset:  position.Offset,
				length:  endPosition.Offset - position.Offset,
				newText: e.newText,
			},
		)
	}

	return result
}
