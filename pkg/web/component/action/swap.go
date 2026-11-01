package action

func (a *Action) Swap(
	target string,
	style string,
) *Action {
	a.target = target
	a.swap = style

	return a
}
