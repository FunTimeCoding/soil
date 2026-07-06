package system

import "github.com/funtimecoding/soil/pkg/constant"

func MarkupFiles(root string) []string {
	return FilesByExtension(
		root,
		constant.MarkupExtension,
		constant.ShortMarkupExtension,
	)
}
