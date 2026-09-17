package pointer_tester

import "github.com/funtimecoding/soil/pkg/lint/pointer"

func Resolver(existing ...string) *pointer.Resolver {
	r := pointer.New()
	r.Roots = Roots()
	r.Exists = exists(existing)
	r.SiblingExists = exists(existing)
	r.PrefixExists = prefixExists(existing)

	return r
}
