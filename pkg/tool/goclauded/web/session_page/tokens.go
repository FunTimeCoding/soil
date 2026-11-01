package session_page

import (
	"github.com/funtimecoding/soil/pkg/generative/anthropic/claude/pricing"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/web/component/table"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"maps"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"slices"
)

func (p *Page) tokens() gomponents.Node {
	usage := p.view.Detail.Usage

	if len(usage) == 0 {
		return gomponents.Group(nil)
	}

	var rows []gomponents.Node

	for _, model := range slices.Sorted(maps.Keys(usage)) {
		t := usage[model]
		name := model

		if !pricing.KnownModel(model) {
			name = join.Empty(model, " (unknown model, sonnet rates)")
		}

		rows = append(
			rows,
			html.Tr(
				html.Td(gomponents.Text(name)),
				numeric(t.Count),
				tokenCell(t.Input),
				tokenCell(t.Output),
				tokenCell(t.CacheCreation),
				tokenCell(t.CacheRead),
				html.Td(
					html.Class(webConstant.NumericClass),
					gomponents.Textf("$%.2f", pricing.Cost(model, t)),
				),
			),
		)
	}

	return gomponents.Group(
		[]gomponents.Node{
			html.H4(gomponents.Text("Token Usage")),
			table.New(
				"Model",
				"Calls",
				"Input",
				"Output",
				"Cache Write",
				"Cache Read",
				"Cost",
			).
				Numeric(1, 2, 3, 4, 5, 6).
				Rows(rows).
				Node(),
		},
	)
}
