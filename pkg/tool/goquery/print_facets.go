package goquery

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/client"
	"sort"
)

func printFacets(facets *[]client.Facet) {
	if facets == nil || len(*facets) == 0 {
		return
	}

	console.Line()

	for _, f := range *facets {
		if f.Values != nil && len(*f.Values) > 0 {
			keys := make([]string, 0, len(*f.Values))

			for k := range *f.Values {
				keys = append(keys, k)
			}

			sort.Strings(keys)

			for _, k := range keys {
				console.Format("  %s=%s (%d)\n", f.Key, k, (*f.Values)[k])
			}
		} else {
			console.Format("  %s: %d distinct\n", f.Key, f.Distinct)
		}
	}
}
