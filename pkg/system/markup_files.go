package system

import soil "github.com/funtimecoding/soil/pkg/constant"

func MarkupFiles(root string) []string {
	return FilesByExtension(
		root,
		soil.MarkupExtension,
		soil.ShortMarkupExtension,
	)
}
