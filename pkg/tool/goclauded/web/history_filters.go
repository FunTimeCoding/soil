package web

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/types/history_filter"
	"github.com/funtimecoding/soil/pkg/web/extended"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func historyFilters(activeKinds []string) gomponents.Node {
	active := map[string]bool{}

	for _, k := range activeKinds {
		active[k] = true
	}

	filters := []*history_filter.Filter{
		history_filter.New("Completions", constant.Complete),
		history_filter.New("Summaries", constant.Summarize),
		history_filter.New("Moments", constant.Moment),
		history_filter.New("Updates", constant.Update),
	}
	var items []gomponents.Node

	for _, f := range filters {
		attrs := []gomponents.Node{
			html.Type("checkbox"),
			html.Name(constant.Kind),
			html.Value(f.Kind),
			extended.Get(constant.HistoryPath),
			extended.Include("#history-filters"),
			extended.Target("#history-content"),
			extended.Swap("innerHTML"),
		}

		if active[f.Kind] {
			attrs = append(attrs, gomponents.Attr("checked", ""))
		}

		items = append(
			items,
			html.Label(html.Input(attrs...), gomponents.Text(f.Label)),
		)
	}

	return html.Form(html.ID("history-filters"), gomponents.Group(items))
}
