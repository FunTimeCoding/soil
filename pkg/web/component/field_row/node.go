package field_row

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func (r *FieldRow) Node() gomponents.Node {
	class := constant.FieldRowClass
	var value gomponents.Node

	switch {
	case len(r.valueNode) > 0:
		value = gomponents.Group(r.valueNode)
	case r.note != "":
		value = html.Small(
			html.Class(constant.FieldNoteClass),
			gomponents.Text(r.note),
		)
	case r.value == "":
		class = join.Space(class, constant.FieldUnsetClass)
		value = html.Span(
			html.Class(constant.AbsentClass),
			gomponents.Text(constant.FieldNotSet),
		)
	default:
		value = gomponents.Text(r.value)
	}

	return html.Div(
		gomponents.If(r.identifier != "", html.ID(r.identifier)),
		html.Class(class),
		gomponents.If(
			r.label != "",
			html.Span(
				html.Class(constant.FieldLabelClass),
				gomponents.If(
					r.tooltip != "",
					gomponents.Attr(constant.TooltipAttribute, r.tooltip),
				),
				gomponents.Text(r.label),
			),
		),
		html.Span(html.Class(constant.FieldValueClass), value),
		gomponents.If(
			len(r.actions) > 0,
			html.Span(
				html.Class(constant.FieldActionsClass),
				gomponents.Group(r.actions),
			),
		),
	)
}
