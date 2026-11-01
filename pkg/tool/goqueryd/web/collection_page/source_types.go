package collection_page

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store/result"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"maps"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"slices"
)

func SourceTypes(
	name string,
	facets []result.Facet,
	current string,
) gomponents.Node {
	index := slices.IndexFunc(
		facets,
		func(f result.Facet) bool { return f.Key == constant.SourceType },
	)

	if index < 0 || len(facets[index].Values) < 2 {
		return gomponents.Group(nil)
	}

	values := facets[index].Values
	chips := []gomponents.Node{sourceChip("All", name, "", current)}

	for _, value := range slices.Sorted(maps.Keys(values)) {
		chips = append(
			chips,
			sourceChip(
				fmt.Sprintf("%s (%d)", value, values[value]),
				name,
				value,
				current,
			),
		)
	}

	return html.Div(
		html.Class(webConstant.ClusterClass),
		gomponents.Group(chips),
	)
}
