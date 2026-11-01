package action

func (a *Action) ButtonClass(class string) *Action {
	a.buttonClass = class

	return a
}
