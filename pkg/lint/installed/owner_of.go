package installed

import "path/filepath"

func OwnerOf(
	root string,
	name string,
) string {
	sources := modules(filepath.Dir(root))

	return owner(sources, name, commands(sources)[name])
}
