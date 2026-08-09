package web

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"net/url"
)

func (s *Server) relationSummary() gomponents.Node {
	relations, e := s.service.ListRelations()

	if e != nil || len(relations) == 0 {
		return nil
	}

	hidden := s.hiddenIdentifiers()
	untyped := 0
	visible := relations[:0]

	for _, relation := range relations {
		if hidden[relation.SourceIdentifier] ||
			hidden[relation.TargetIdentifier] {
			continue
		}

		visible = append(visible, relation)

		if relation.Type == "" {
			untyped++
		}
	}

	relations = visible

	if len(relations) == 0 {
		return nil
	}

	entries := []gomponents.Node{
		html.A(
			gomponents.Attr("href", constant.RelationsPath),
			gomponents.Textf("%d relations", len(relations)),
		),
	}

	if untyped > 0 {
		params := url.Values{}
		params.Set(constant.Type, constant.UntypedFilter)
		entries = append(
			entries,
			gomponents.Text(" · "),
			html.A(
				gomponents.Attr(
					"href",
					fmt.Sprintf(
						"%s?%s",
						constant.RelationsPath,
						params.Encode(),
					),
				),
				gomponents.Textf("%d untyped", untyped),
			),
		)
	}

	return html.P(entries...)
}
