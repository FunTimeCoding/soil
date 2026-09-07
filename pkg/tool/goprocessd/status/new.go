package status

func New(
	name string,
	command string,
	running bool,
	startedAt string,
) *Status {
	return &Status{
		Name:      name,
		Command:   command,
		Running:   running,
		StartedAt: startedAt,
	}
}
