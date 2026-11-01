package search_page

import (
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store/result"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func facets(
	all []result.Facet,
	query string,
	collection string,
	metadata map[string]string,
) gomponents.Node {
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

		sections = append(sections, facet(f, query, collection, metadata))
	}

	return gomponents.Group(sections)
}
