package field

import (
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func (f *Field) suggestionNode() gomponents.Node {
	if len(f.suggestions) == 0 {
		return gomponents.Group(nil)
	}

	return html.DataList(
		html.ID(f.suggestionList()),
		gomponents.Map(
			f.suggestions,
			func(v string) gomponents.Node { return html.Option(html.Value(v)) },
		),
	)
}
