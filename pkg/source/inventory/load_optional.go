package inventory

import "github.com/funtimecoding/soil/pkg/system"

func LoadOptional(path string) *Inventory {
	if !system.FileExists(path) {
		return New()
	}

	return Load(path)
}
