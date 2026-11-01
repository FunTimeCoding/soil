package action

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/extended"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"strings"
)

func (a *Action) Node() gomponents.Node {
	classes := []string{constant.ActionClass}

	if a.small {
		classes = append(classes, constant.ActionSmallClass)
	}

	if a.class != "" {
		classes = append(classes, a.class)
	}

	form := []gomponents.Node{
		html.Class(strings.Join(classes, " ")),
		html.Method("post"),
		html.Action(a.path),
	}

	if a.target != "" {
		form = append(
			form,
			extended.Post(a.path),
			extended.Target(a.target),
			extended.Swap(a.swap),
		)
	}

	if a.confirm != "" && a.target != "" {
		form = append(form, extended.Confirm(a.confirm))
	}

	if a.confirm != "" && a.target == "" {
		form = append(
			form,
			gomponents.Attr(
				"onsubmit",
				fmt.Sprintf("return confirm(%q)", a.confirm),
			),
		)
	}

	form = append(form, a.fields...)

	return html.Form(
		append(
			form,
			html.Button(
				html.Type(constant.FormSubmit),
				gomponents.If(a.buttonClass != "", html.Class(a.buttonClass)),
				gomponents.If(a.title != "", html.Title(a.title)),
				gomponents.If(a.disabled, html.Disabled()),
				gomponents.If(a.label != "", gomponents.Text(a.label)),
			),
		)...,
	)
}
