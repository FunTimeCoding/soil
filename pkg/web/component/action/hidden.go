package action

import "maragu.dev/gomponents/html"

func (a *Action) Hidden(
	name string,
	value string,
) *Action {
	a.fields = append(
		a.fields,
		html.Input(html.Type("hidden"), html.Name(name), html.Value(value)),
	)

	return a
}
