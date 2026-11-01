package action

func (a *Action) Title(title string) *Action {
	a.title = title

	return a
}
