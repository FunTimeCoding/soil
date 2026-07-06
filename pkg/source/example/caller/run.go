package caller

import "github.com/funtimecoding/soil/pkg/source/example/target"

func Run(name string) string {
	if target.IsValid(name) {
		return target.Check(name)
	}

	return ""
}
