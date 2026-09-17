package pointer_tester

import (
	"github.com/funtimecoding/soil/pkg/lint"
	"github.com/funtimecoding/soil/pkg/lint/pointer"
)

func Literals(
	contents map[string][]string,
	registries []string,
	implicitBases []string,
	existing ...string,
) lint.Checker {
	r := Resolver(existing...)
	r.ImplicitBases = implicitBases
	r.Registries = registries
	r.Literal = func(
		directory string,
		needle string,
	) bool {
		for _, content := range contents[directory] {
			if pointer.ContainsLiteral(content, needle) {
				return true
			}
		}

		return false
	}

	return lint.Pointers(r, discard)
}
