package web

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/strings/split"
)

func shortenPaths(files string) string {
	var result []string

	for _, line := range split.NewLine(files) {
		result = append(result, shortenPath(line))
	}

	return join.CommaSpace(result)
}
