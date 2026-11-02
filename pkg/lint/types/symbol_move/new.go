package symbol_move

func New(
	directory string,
	name string,
	newDirectory string,
	newName string,
) *Move {
	return &Move{
		Directory:    directory,
		Name:         name,
		NewDirectory: newDirectory,
		NewName:      newName,
	}
}
