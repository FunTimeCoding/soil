package result

func NewFileReferences(
	file string,
	symbols []*References,
) *FileReferences {
	return &FileReferences{
		File:    file,
		Symbols: symbols,
	}
}
