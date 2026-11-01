package entry

import (
	"github.com/funtimecoding/soil/pkg/web/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func (e *Entry) Alternative(
	path string,
	label string,
) *Entry {
	e.alternatives = append(
		e.alternatives,
		html.Button(
			html.Type(constant.FormSubmit),
			html.Class(constant.EntryAlternativeClass),
			html.FormAction(path),
			gomponents.Attr("formnovalidate"),
			gomponents.Text(label),
		),
	)

	return e
}
