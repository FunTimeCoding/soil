package action

func (a *Action) Disabled(disabled bool) *Action {
	a.disabled = disabled

	return a
}
