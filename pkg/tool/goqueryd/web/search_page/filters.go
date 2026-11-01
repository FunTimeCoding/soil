package search_page

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/web/component/chip"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"maps"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"slices"
)

func filters(
	query string,
	collection string,
	metadata map[string]string,
) gomponents.Node {
	if len(metadata) == 0 {
		return gomponents.Group(nil)
	}

	var chips []gomponents.Node

	for _, k := range slices.Sorted(maps.Keys(metadata)) {
		chips = append(
			chips,
			chip.New(fmt.Sprintf("%s: %s", k, metadata[k])).
				Link(Path(query, collection, copyWithout(metadata, k))).
				Content(gomponents.Text(" ×")).
				Node(),
		)
	}

	return html.Div(
		html.Class(webConstant.ClusterClass),
		gomponents.Text("Filtered:"),
		gomponents.Group(chips),
	)
}
