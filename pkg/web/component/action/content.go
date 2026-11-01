package action

import "maragu.dev/gomponents"

func (a *Action) Content(n ...gomponents.Node) *Action {
	a.fields = append(a.fields, n...)

	return a
}
