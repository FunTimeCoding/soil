package command

type CommandError struct {
	Command string
	Output  string
	Stderr  string
	Wrapped error
}
