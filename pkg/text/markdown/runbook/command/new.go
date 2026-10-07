package command

func New(
	description string,
	code string,
) *Command {
	return &Command{Description: description, Code: code}
}
