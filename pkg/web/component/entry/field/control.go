package field

import (
	"github.com/funtimecoding/soil/pkg/web/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func (f *Field) control() gomponents.Node {
	switch f.kind {
	case constant.EntrySelectKind:
		return f.selectNode()
	case constant.EntryTextAreaKind:
		return html.Textarea(
			html.Name(f.name),
			gomponents.If(f.required, html.Required()),
			gomponents.If(f.placeholder != "", html.Placeholder(f.placeholder)),
			gomponents.If(f.lines != "", html.Rows(f.lines)),
			gomponents.Text(f.value),
		)
	default:
		return html.Input(
			html.Type(f.kind),
			html.Name(f.name),
			gomponents.If(f.value != "", html.Value(f.value)),
			gomponents.If(f.placeholder != "", html.Placeholder(f.placeholder)),
			gomponents.If(f.required, html.Required()),
			gomponents.If(f.focus, html.AutoFocus()),
			gomponents.If(f.minimum != "", html.Min(f.minimum)),
			gomponents.If(f.maximum != "", html.Max(f.maximum)),
			gomponents.If(f.step != "", html.Step(f.step)),
			gomponents.If(f.accept != "", html.Accept(f.accept)),
			gomponents.If(
				f.suggestionList() != "",
				html.List(f.suggestionList()),
			),
			gomponents.If(f.class != "", html.Class(f.class)),
			f.revealAttributes(),
		)
	}
}
