package entry

import (
	"github.com/funtimecoding/soil/pkg/web/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func (e *Entry) body() []gomponents.Node {
	var result []gomponents.Node
	var grouped []gomponents.Node
	legend := ""
	open := false
	flush := func() {
		if open {
			result = append(
				result,
				html.FieldSet(
					html.Legend(gomponents.Text(legend)),
					gomponents.Group(grouped),
				),
			)
		}
	}

	for _, f := range e.fields {
		if f.Kind() == constant.EntryGroupKind {
			flush()
			legend, grouped, open = f.Label(), nil, f.Label() != ""

			continue
		}

		if open {
			grouped = append(grouped, f.Node())
		} else {
			result = append(result, f.Node())
		}
	}

	flush()

	return result
}
