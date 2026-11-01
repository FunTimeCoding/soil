package search_page

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store/result"
	"github.com/funtimecoding/soil/pkg/web/component/chip"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"maps"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"slices"
)

func facet(
	f result.Facet,
	query string,
	collection string,
	metadata map[string]string,
) gomponents.Node {
	chips := make([]gomponents.Node, 0, len(f.Values))

	for _, value := range slices.Sorted(maps.Keys(f.Values)) {
		c := chip.New(fmt.Sprintf("%s (%d)", value, f.Values[value]))

		if metadata[f.Key] == value {
			c.Link(Path(query, collection, copyWithout(metadata, f.Key))).
				Current()
		} else {
			c.Link(Path(query, collection, copyWith(metadata, f.Key, value)))
		}

		chips = append(chips, c.Node())
	}

	return html.Div(
		html.Class(webConstant.ClusterClass),
		gomponents.Text(f.Key),
		gomponents.Group(chips),
	)
}
