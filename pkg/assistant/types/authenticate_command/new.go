package authenticate_command

func New(
	typeValue string,
	token string,
) *Command {
	return &Command{Type: typeValue, Token: token}
}
