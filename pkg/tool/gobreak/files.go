package gobreak

import (
	"github.com/funtimecoding/soil/pkg/tool/gobreak/break_file"
	"slices"
)

func files(b *break_file.Break) []string {
	var result []string

	for _, e := range b.Edits {
		if !slices.Contains(result, e.File) {
			result = append(result, e.File)
		}
	}

	return result
}
