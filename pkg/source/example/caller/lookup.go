package caller

import "github.com/funtimecoding/soil/pkg/source/example/target"

func Lookup(
	e *target.Entry,
	name string,
) bool {
	return e.FindByName(name)
}
