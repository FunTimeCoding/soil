package symbol

func New(
	directory string,
	clause string,
	receiver string,
	name string,
	kind string,
	file string,
	line int,
) *Symbol {
	return &Symbol{
		Package:  directory,
		Clause:   clause,
		Receiver: receiver,
		Name:     name,
		Kind:     kind,
		File:     file,
		Line:     line,
	}
}
