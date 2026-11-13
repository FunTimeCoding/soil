package gofix

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"slices"
)

func loadKey(
	directory string,
	patterns []string,
) string {
	return join.NewLine(slices.Concat([]string{directory}, patterns))
}
