package pointer

type Resolver struct {
	Roots         []string
	ImplicitBases []string
	Registries    []string
	Exists        func(string) bool
	SiblingExists func(string) bool
	Ignored       func(string) bool
	Stdlib        func(string) bool
	Dependency    func(string) bool
	PrefixExists  func(string, string) bool
	Literal       func(string, string) bool
	Routes        func(string) ([]string, bool)
}
