package web

import (
	"github.com/funtimecoding/soil/pkg/generative/anthropic/claude/pricing"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"sort"
)

func usageTable(usage map[string]*pricing.Tokens) gomponents.Node {
	models := make([]string, 0, len(usage))

	for model := range usage {
		models = append(models, model)
	}

	sort.Strings(models)
	var rows []gomponents.Node

	for _, model := range models {
		t := usage[model]
		name := model

		if !pricing.KnownModel(model) {
			name = join.Empty(model, " (unknown model, sonnet rates)")
		}

		rows = append(
			rows,
			html.Tr(
				html.Td(gomponents.Text(name)),
				html.Td(gomponents.Textf("%d", t.Count)),
				html.Td(gomponents.Text(pricing.FormatTokens(t.Input))),
				html.Td(gomponents.Text(pricing.FormatTokens(t.Output))),
				html.Td(gomponents.Text(pricing.FormatTokens(t.CacheCreation))),
				html.Td(gomponents.Text(pricing.FormatTokens(t.CacheRead))),
				html.Td(gomponents.Textf(
					"$%.2f",
					pricing.Cost(model, t),
				)),
			),
		)
	}

	return html.Table(
		html.THead(
			html.Tr(
				html.Th(gomponents.Text("Model")),
				html.Th(gomponents.Text("Calls")),
				html.Th(gomponents.Text("Input")),
				html.Th(gomponents.Text("Output")),
				html.Th(gomponents.Text("Cache Write")),
				html.Th(gomponents.Text("Cache Read")),
				html.Th(gomponents.Text("Cost")),
			),
		),
		html.TBody(rows...),
	)
}
