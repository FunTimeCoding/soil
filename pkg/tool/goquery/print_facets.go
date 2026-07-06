package goquery

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/client"
	"sort"
)

func printFacets(facets *[]client.Facet) {
	if facets == nil || len(*facets) == 0 {
		return
	}

	fmt.Println()

	for _, f := range *facets {
		if f.Values != nil && len(*f.Values) > 0 {
			keys := make([]string, 0, len(*f.Values))

			for k := range *f.Values {
				keys = append(keys, k)
			}

			sort.Strings(keys)

			for _, k := range keys {
				fmt.Printf("  %s=%s (%d)\n", f.Key, k, (*f.Values)[k])
			}
		} else {
			fmt.Printf("  %s: %d distinct\n", f.Key, f.Distinct)
		}
	}
}
