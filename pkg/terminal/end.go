package terminal

func (t *Terminal) end(
	outcome string,
	code int,
) {
	t.ender.EndCommand(outcome)
	t.exit(code)
}
