package mutation

func NewMutant(
	status string,
	line int,
) *Mutant {
	return &Mutant{Status: status, Line: line}
}
