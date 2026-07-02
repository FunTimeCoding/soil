package caller

import "github.com/funtimecoding/go-library/pkg/source/example/target"

func Lookup(
	e *target.Entry,
	name string,
) bool {
	return e.FindByName(name)
}
