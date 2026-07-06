package client

import (
	"github.com/funtimecoding/soil/pkg/kubernetes/client/client_configuration"
	"slices"
	"sort"
	"strings"
)

func (c *Client) ConfigurationContexts(filterExtended bool) []string {
	var result []string

	for context := range client_configuration.Interface().Contexts {
		result = append(result, context)
	}

	if filterExtended {
		var filtered []string

		for _, e := range result {
			if slices.Contains(filtered, e) {
				continue
			}

			var isExtendedName bool

			for _, inner := range result {
				if e == inner {
					continue
				}

				if strings.HasPrefix(e, inner) {
					isExtendedName = true
				}
			}

			if isExtendedName {
				continue
			}

			filtered = append(filtered, e)
		}

		result = filtered
	}

	sort.Strings(result)

	return result
}
