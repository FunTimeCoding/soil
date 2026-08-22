package command

func New(
	name string,
	output string,
	stderr string,
	wrapped error,
) *CommandError {
	return &CommandError{
		Command: name,
		Output:  output,
		Stderr:  stderr,
		Wrapped: wrapped,
	}
}
