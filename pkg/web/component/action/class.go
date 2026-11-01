package action

func (a *Action) Class(class string) *Action {
	a.class = class

	return a
}
