package pointer

func New() *Resolver {
	return &Resolver{
		Exists:        absent,
		SiblingExists: absent,
		Ignored:       absent,
		Stdlib:        absent,
		Dependency:    absent,
		PrefixExists:  absentIn,
		Literal:       absentIn,
		Routes:        noRoutes,
	}
}
