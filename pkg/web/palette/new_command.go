package palette

func NewCommand(
	label string,
	path string,
	category string,
) *Command {
	return &Command{Label: label, Path: path, Category: category}
}
