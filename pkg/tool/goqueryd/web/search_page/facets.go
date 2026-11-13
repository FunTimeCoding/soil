package search_page

import (
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store/result"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func (p *Page) facets(all []result.Facet) gomponents.Node {
	sections := make([]gomponents.Node, 0, len(all))

	for _, f := range all {
		if len(f.Values) == 0 {
			sections = append(
				sections,
				html.P(
					html.Small(
						gomponents.Textf("%s: %d distinct", f.Key, f.Distinct),
					),
				),
			)

			continue
		}

		sections = append(sections, p.facet(f))
	}

	return gomponents.Group(sections)
}
