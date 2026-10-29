package service

import (
	"github.com/funtimecoding/soil/pkg/system"
	"strings"
)

func countTiddlers(workdir string) int {
	var result int

	for _, entry := range system.ReadDirectory(workdir) {
		if strings.HasSuffix(entry.Name(), ".tid") {
			result++
		}
	}

	return result
}
