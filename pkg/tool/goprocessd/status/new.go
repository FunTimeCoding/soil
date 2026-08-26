package status

func New(
	name string,
	command string,
	running bool,
) *Status {
	return &Status{Name: name, Command: command, Running: running}
}
