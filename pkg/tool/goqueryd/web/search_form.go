package web

import (
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store/status"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func searchForm(
	query string,
	collection string,
	collections []status.CollectionStatus,
) gomponents.Node {
	var options []gomponents.Node
	allAttrs := []gomponents.Node{html.Value("")}

	if collection == "" {
		allAttrs = append(allAttrs, gomponents.Attr("selected", ""))
	}

	options = append(
		options,
		html.Option(
			gomponents.Text("All collections"),
			gomponents.Group(allAttrs),
		),
	)

	for _, c := range collections {
		attrs := []gomponents.Node{html.Value(c.Name)}

		if collection == c.Name {
			attrs = append(attrs, gomponents.Attr("selected", ""))
		}

		options = append(
			options,
			html.Option(
				gomponents.Text(c.Name),
				gomponents.Group(attrs),
			),
		)
	}

	return html.Form(
		gomponents.Attr("method", "GET"),
		gomponents.Attr("action", constant.SearchPath),
		html.Input(
			html.Type(constant.Search),
			html.Name("query"),
			gomponents.Attr("placeholder", "search documents..."),
			gomponents.Attr("value", query),
			gomponents.Attr("autocomplete", "off"),
		),
		html.Select(
			html.Name(constant.Collection),
			gomponents.Group(options),
		),
		html.Button(
			html.Type("submit"),
			gomponents.Text(constant.SearchTitle),
		),
	)
}
