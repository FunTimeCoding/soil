package system

import soil "github.com/funtimecoding/soil/pkg/constant"

func GoFiles(root string) []string {
	return FilesByExtension(root, soil.GoExtension)
}
