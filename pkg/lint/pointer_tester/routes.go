package pointer_tester

import "github.com/funtimecoding/soil/pkg/lint"

func Routes(
	specifications map[string][]string,
	existing ...string,
) lint.Checker {
	r := Resolver(existing...)
	r.Routes = func(directory string) ([]string, bool) {
		paths, found := specifications[directory]

		return paths, found
	}

	return lint.Pointers(r, discard)
}
