package entry

import (
	"github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/layout"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"strings"
)

func (e *Entry) Node() gomponents.Node {
	classes := []string{constant.EntryClass}

	if e.inline {
		classes = append(classes, constant.EntryInlineClass)
	}

	if e.record {
		classes = append(classes, constant.EntryRecordClass)
	}

	if e.class != "" {
		classes = append(classes, e.class)
	}

	button := html.Button(
		html.Type(constant.FormSubmit),
		gomponents.If(e.buttonClass != "", html.Class(e.buttonClass)),
		gomponents.Text(e.submit),
	)

	if len(e.alternatives) > 0 || len(e.beside) > 0 {
		button = html.Div(
			html.Class(constant.EntryButtonsClass),
			button,
			gomponents.Group(e.alternatives),
			gomponents.Group(e.beside),
		)
	}

	return html.Form(
		html.Class(strings.Join(classes, " ")),
		gomponents.If(e.identifier != "", html.ID(e.identifier)),
		html.Method(e.method()),
		html.Action(e.path),
		gomponents.If(e.carriesFile(), html.EncType(constant.Form)),
		gomponents.Group(e.extendedAttributes()),
		gomponents.If(e.heading != "", html.H5(gomponents.Text(e.heading))),
		gomponents.If(e.alert != "", layout.Alert(e.alert)),
		gomponents.Group(e.hidden),
		gomponents.Group(e.body()),
		button,
	)
}
