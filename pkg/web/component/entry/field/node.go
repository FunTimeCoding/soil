package field

import (
	"github.com/funtimecoding/soil/pkg/web/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func (f *Field) Node() gomponents.Node {
	if f.kind == constant.EntryContentKind {
		return gomponents.Group(f.content)
	}

	if f.kind == constant.EntryCheckboxKind {
		return html.Label(
			html.Input(
				html.Type(f.kind),
				html.Name(f.name),
				html.Value(f.value),
				gomponents.If(f.checked, html.Checked()),
			),
			gomponents.Text(f.label),
		)
	}

	return html.Label(
		html.Span(
			html.Class(constant.EntryLabelClass),
			gomponents.Text(f.label),
		),
		f.control(),
		f.suggestionNode(),
	)
}
