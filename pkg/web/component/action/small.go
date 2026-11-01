package action

func (a *Action) Small() *Action {
	a.small = true

	return a
}
