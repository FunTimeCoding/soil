package web

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func historyFilters(activeKinds []string) gomponents.Node {
	active := map[string]bool{}

	for _, k := range activeKinds {
		active[k] = true
	}

	filters := []historyFilter{
		{"Completions", constant.Complete},
		{"Summaries", constant.Summarize},
		{"Moments", constant.Moment},
		{"Updates", constant.Update},
	}
	var items []gomponents.Node

	for _, f := range filters {
		attrs := []gomponents.Node{
			html.Type("checkbox"),
			html.Name(constant.Kind),
			html.Value(f.kind),
			gomponents.Attr("hx-get", constant.HistoryPath),
			gomponents.Attr("hx-include", "#history-filters"),
			gomponents.Attr("hx-target", "#history-content"),
			gomponents.Attr("hx-swap", "innerHTML"),
		}

		if active[f.kind] {
			attrs = append(attrs, gomponents.Attr("checked", ""))
		}

		items = append(
			items,
			html.Label(
				html.Input(attrs...),
				gomponents.Text(f.label),
			),
		)
	}

	return html.Form(
		html.ID("history-filters"),
		gomponents.Group(items),
	)
}
