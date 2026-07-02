package caller

import "github.com/funtimecoding/go-library/pkg/source/example/target"

func Run(name string) string {
	if target.IsValid(name) {
		return target.Check(name)
	}

	return ""
}
