package action

func (a *Action) Confirm(message string) *Action {
	a.confirm = message

	return a
}
