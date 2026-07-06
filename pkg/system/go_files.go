package system

import "github.com/funtimecoding/soil/pkg/constant"

func GoFiles(root string) []string {
	return FilesByExtension(root, constant.GoExtension)
}
