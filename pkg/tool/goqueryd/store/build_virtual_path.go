package store

import (
	"github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
)

func buildVirtualPath(
	collection string,
	path string,
) string {
	return join.Empty("qmd://", collection, constant.Slash, path)
}
