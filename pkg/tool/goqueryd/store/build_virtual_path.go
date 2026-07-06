package store

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/strings/separator"
)

func buildVirtualPath(
	collection string,
	path string,
) string {
	return join.Empty("qmd://", collection, separator.Slash, path)
}
