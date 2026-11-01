package field

import (
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func (f *Field) selectNode() gomponents.Node {
	options := make([]gomponents.Node, 0, len(f.choices)+1)

	if f.prompt != "" {
		options = append(
			options,
			html.Option(
				html.Value(""),
				gomponents.If(f.required, html.Disabled()),
				gomponents.If(f.value == "", html.Selected()),
				gomponents.Text(f.prompt),
			),
		)
	}

	for _, c := range f.choices {
		options = append(
			options,
			html.Option(
				html.Value(c.Value),
				gomponents.If(c.Value == f.value, html.Selected()),
				gomponents.Text(c.Label),
			),
		)
	}

	return html.Select(
		html.Name(f.name),
		gomponents.If(f.required, html.Required()),
		f.revealAttributes(),
		gomponents.Group(options),
	)
}
