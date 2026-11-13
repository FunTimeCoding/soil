package lint

import (
	"fmt"
	"path/filepath"
)

func HeaderLine(
	name string,
	root string,
	detail string,
) string {
	if detail == "" {
		return fmt.Sprintf("%s: %s (%s)", name, filepath.Base(root), root)
	}

	return fmt.Sprintf(
		"%s: %s (%s) %s",
		name,
		filepath.Base(root),
		root,
		detail,
	)
}
