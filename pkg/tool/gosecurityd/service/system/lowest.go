package system

import "github.com/funtimecoding/soil/pkg/vulnerability/debian/version"

func lowest(
	a string,
	b string,
) string {
	if a == "" || (b != "" && version.Compare(b, a) < 0) {
		return b
	}

	return a
}
