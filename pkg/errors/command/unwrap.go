package command

func (e *CommandError) Unwrap() error {
	return e.Wrapped
}
