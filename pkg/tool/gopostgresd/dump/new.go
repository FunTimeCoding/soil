package dump

func New(program ...string) *Program {
	return &Program{program: program}
}
