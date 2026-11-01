package entry

import "maragu.dev/gomponents/html"

func (e *Entry) Hidden(
	name string,
	value string,
) *Entry {
	e.hidden = append(
		e.hidden,
		html.Input(html.Type("hidden"), html.Name(name), html.Value(value)),
	)

	return e
}
