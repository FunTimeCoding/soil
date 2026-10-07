package pointer_tester

import "github.com/funtimecoding/soil/pkg/lint/pointer/resolver"

func Resolver(existing ...string) *resolver.Resolver {
	r := resolver.New()
	r.Roots = Roots()
	r.Exists = exists(existing)
	r.SiblingExists = exists(existing)
	r.PrefixExists = prefixExists(existing)

	return r
}
