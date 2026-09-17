package lint

import (
	"github.com/funtimecoding/soil/pkg/console"
	"path/filepath"
)

func Header(
	name string,
	root string,
	detail string,
) {
	if detail == "" {
		console.Format("%s: %s (%s)\n", name, filepath.Base(root), root)

		return
	}

	console.Format("%s: %s (%s) %s\n", name, filepath.Base(root), root, detail)
}
