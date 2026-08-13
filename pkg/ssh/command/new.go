package command

func New(command string) *Command {
	return &Command{Command: command, Environment: make(map[string]string)}
}
