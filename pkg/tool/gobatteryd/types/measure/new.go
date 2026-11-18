package measure

func New(
	variable string,
	name string,
	help string,
) *Measure {
	return &Measure{Variable: variable, Name: name, Help: help}
}
