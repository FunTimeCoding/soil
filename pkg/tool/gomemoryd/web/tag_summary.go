package web

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func (s *Server) tagSummary() gomponents.Node {
	tags, e := s.service.ListTags()

	if e != nil {
		return html.P(gomponents.Text("Failed to load tags."))
	}

	if len(tags) == 0 {
		return gomponents.Group(
			[]gomponents.Node{
				html.H3(gomponents.Text("Tags")),
				html.P(gomponents.Text("No tags.")),
			},
		)
	}

	var pips []gomponents.Node

	for _, t := range tags {
		pips = append(
			pips,
			html.A(
				gomponents.Attr(
					"href",
					fmt.Sprintf(
						"%s?%s=%s",
						constant.MemoriesPath,
						constant.Tag,
						t.Tag,
					),
				),
				html.Class("tag-pip"),
				gomponents.Textf("%s (%d)", t.Tag, t.Count),
			),
			gomponents.Text(" "),
		)
	}

	return gomponents.Group(
		[]gomponents.Node{
			html.H3(gomponents.Text("Tags")),
			html.P(gomponents.Group(pips)),
		},
	)
}
