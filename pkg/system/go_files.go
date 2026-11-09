package system

import soil "github.com/funtimecoding/soil/pkg/constant"

func GoFiles(root string) []string {
	return ExtensionFiles(root, soil.GoExtension)
}
