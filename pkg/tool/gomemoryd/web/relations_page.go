package web

import (
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"net/http"
)

func (s *Server) relationsPage(
	w http.ResponseWriter,
	r *http.Request,
) {
	filter := r.URL.Query().Get(constant.Type)
	relations, e := s.service.ListRelations()

	if e != nil {
		s.view.RenderPage(
			w,
			constant.RelationsTitle,
			constant.RelationsPath,
			html.P(gomponents.Text("Failed to load relations.")),
		)

		return
	}

	hidden := s.hiddenIdentifiers()
	visible := relations[:0]

	for _, relation := range relations {
		if hidden[relation.SourceIdentifier] ||
			hidden[relation.TargetIdentifier] {
			continue
		}

		visible = append(visible, relation)
	}

	relations = visible
	untyped := 0
	typeCounts := map[string]int{}

	for _, relation := range relations {
		if relation.Type == "" {
			untyped++

			continue
		}

		typeCounts[relation.Type]++
	}

	total := len(relations)
	var content []gomponents.Node
	content = append(
		content,
		html.H3(gomponents.Text(constant.RelationsTitle)),
		relationTypeFilter(filter, total, untyped, typeCounts),
	)
	var rows []gomponents.Node

	for _, relation := range relations {
		if filter == constant.UntypedFilter && relation.Type != "" {
			continue
		}

		if filter != "" &&
			filter != constant.UntypedFilter &&
			relation.Type != filter {
			continue
		}

		label := relation.Type

		if label == "" {
			label = "-"
		}

		rows = append(
			rows,
			html.Tr(
				html.Td(memoryLink(
					relation.SourceIdentifier,
					relation.SourceName,
				)),
				html.Td(scopeText(relation.SourceScope)),
				html.Td(gomponents.Text(label)),
				html.Td(memoryLink(
					relation.TargetIdentifier,
					relation.TargetName,
				)),
				html.Td(scopeText(relation.TargetScope)),
			),
		)
	}

	if len(rows) == 0 {
		content = append(
			content,
			html.P(gomponents.Text("No relations.")),
		)
	} else {
		content = append(
			content,
			html.Table(
				html.THead(
					html.Tr(
						html.Th(gomponents.Text("Source")),
						html.Th(gomponents.Text("Scope")),
						html.Th(gomponents.Text("Type")),
						html.Th(gomponents.Text("Target")),
						html.Th(gomponents.Text("Scope")),
					),
				),
				html.TBody(rows...),
			),
		)
	}

	s.view.RenderPage(
		w,
		constant.RelationsTitle,
		constant.RelationsPath,
		content...,
	)
}
